package rpcserver

import (
	"context"
	"crypto"
	"crypto/tls"
	"net"
	"net/http"
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"go.f110.dev/heimdallr/pkg/auth"
	"go.f110.dev/heimdallr/pkg/cert"
	"go.f110.dev/heimdallr/pkg/config/configv2"
	"go.f110.dev/heimdallr/pkg/database"
	"go.f110.dev/heimdallr/pkg/logger"
	"go.f110.dev/heimdallr/pkg/rpc"
	"go.f110.dev/heimdallr/pkg/rpc/rpcservice"
)

var (
	registry     = prometheus.NewRegistry()
	registerOnce sync.Once
)

type Server struct {
	Config *configv2.Config

	ca            *cert.CertificateAuthority
	server        *grpc.Server
	privKey       crypto.PrivateKey
	serverMetrics *grpc_prometheus.ServerMetrics
}

func unaryAccessLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Log.Debug("Unary", zap.String("method", info.FullMethod))
	return handler(ctx, req)
}

func streamAccessLogInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger.Log.Debug("Stream", zap.String("method", info.FullMethod))
	return handler(srv, ss)
}

func NewServer(
	conf *configv2.Config,
	user database.UserDatabase,
	token database.TokenDatabase,
	cluster database.ClusterDatabase,
	relay database.RelayLocator,
	ca *cert.CertificateAuthority,
	isReady func() bool,
) *Server {
	r := grpc_prometheus.NewServerMetrics()
	grpc_zap.ReplaceGrpcLoggerV2(logger.Log)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryAccessLogInterceptor,
			auth.UnaryInterceptor,
			r.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamAccessLogInterceptor,
			auth.StreamInterceptor,
			r.StreamServerInterceptor(),
		)),
	)
	rpc.RegisterClusterServer(s, rpcservice.NewClusterService(user, token, cluster, relay))
	rpc.RegisterAdminServer(s, rpcservice.NewAdminService(conf, user))
	rpc.RegisterCertificateAuthorityServer(s, rpcservice.NewCertificateAuthorityService(conf, ca))
	rpc.RegisterUserServer(s, rpcservice.NewUserService(user))
	healthpb.RegisterHealthServer(s, rpcservice.NewHealthService(isReady))
	r.InitializeMetrics(s)
	registerOnce.Do(func() {
		registry.MustRegister(r)
		registry.MustRegister(prometheus.NewGoCollector())
	})

	return &Server{
		Config:        conf,
		ca:            ca,
		server:        s,
		serverMetrics: r,
	}
}

func (s *Server) Start() error {
	logger.Log.Info("Start RPC server", zap.String("listen", s.Config.RPCServer.Bind), zap.String("hostname", rpc.ServerHostname))
	serverCert, privKey, err := s.ca.NewServerCertificate(rpc.ServerHostname)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	l, err := net.Listen("tcp", s.Config.RPCServer.Bind)
	if err != nil {
		return xerrors.Errorf(": %v", err)
	}

	tlsCert := tls.Certificate{
		Certificate: [][]byte{serverCert.Raw},
		PrivateKey:  privKey,
	}
	listener := tls.NewListener(l, &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	})

	go func() {
		if s.Config.RPCServer.MetricsBind == "" {
			return
		}

		handler := promhttp.InstrumentMetricHandler(registry, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
		mux := http.NewServeMux()
		mux.Handle("/metrics", handler)
		logger.Log.Info("Start RPC metrics server", zap.String("listen", s.Config.RPCServer.MetricsBind))
		http.ListenAndServe(s.Config.RPCServer.MetricsBind, mux)
	}()

	return s.server.Serve(listener)
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.Log.Info("Shutdown RPC server")
	s.server.GracefulStop()
	return nil
}
