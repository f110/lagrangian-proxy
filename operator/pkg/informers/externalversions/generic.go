/*

MIT License

Copyright (c) 2019 Fumihiro Ito

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	v1alpha1 "go.f110.dev/heimdallr/operator/pkg/api/etcd/v1alpha1"
	proxyv1alpha1 "go.f110.dev/heimdallr/operator/pkg/api/proxy/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=cert-manager.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("certificates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha2().Certificates().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("certificaterequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha2().CertificateRequests().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("clusterissuers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha2().ClusterIssuers().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("issuers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha2().Issuers().Informer()}, nil

		// Group=etcd.f110.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("etcdclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Etcd().V1alpha1().EtcdClusters().Informer()}, nil

		// Group=monitoring.coreos.com, Version=v1
	case v1.SchemeGroupVersion.WithResource("alertmanagers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().Alertmanagers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("podmonitors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().PodMonitors().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("prometheuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().Prometheuses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("prometheusrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().PrometheusRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("servicemonitors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().ServiceMonitors().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("thanosrulers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().ThanosRulers().Informer()}, nil

		// Group=proxy.f110.dev, Version=v1alpha1
	case proxyv1alpha1.SchemeGroupVersion.WithResource("backends"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Proxy().V1alpha1().Backends().Informer()}, nil
	case proxyv1alpha1.SchemeGroupVersion.WithResource("proxies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Proxy().V1alpha1().Proxies().Informer()}, nil
	case proxyv1alpha1.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Proxy().V1alpha1().Roles().Informer()}, nil
	case proxyv1alpha1.SchemeGroupVersion.WithResource("rolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Proxy().V1alpha1().RoleBindings().Informer()}, nil
	case proxyv1alpha1.SchemeGroupVersion.WithResource("rpcpermissions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Proxy().V1alpha1().RpcPermissions().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
