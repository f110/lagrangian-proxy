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

package v1

import (
	"context"
	time "time"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	versioned "go.f110.dev/heimdallr/operator/pkg/client/versioned"
	internalinterfaces "go.f110.dev/heimdallr/operator/pkg/informers/externalversions/internalinterfaces"
	v1 "go.f110.dev/heimdallr/operator/pkg/listers/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceMonitorInformer provides access to a shared informer and lister for
// ServiceMonitors.
type ServiceMonitorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceMonitorLister
}

type serviceMonitorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceMonitorInformer constructs a new informer for ServiceMonitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceMonitorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceMonitorInformer constructs a new informer for ServiceMonitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().ServiceMonitors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().ServiceMonitors(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1.ServiceMonitor{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceMonitorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceMonitorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceMonitorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1.ServiceMonitor{}, f.defaultInformer)
}

func (f *serviceMonitorInformer) Lister() v1.ServiceMonitorLister {
	return v1.NewServiceMonitorLister(f.Informer().GetIndexer())
}
