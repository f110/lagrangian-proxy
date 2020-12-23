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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "go.f110.dev/heimdallr/operator/pkg/api/etcd/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EtcdClusterLister helps list EtcdClusters.
// All objects returned here must be treated as read-only.
type EtcdClusterLister interface {
	// List lists all EtcdClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.EtcdCluster, err error)
	// EtcdClusters returns an object that can list and get EtcdClusters.
	EtcdClusters(namespace string) EtcdClusterNamespaceLister
	EtcdClusterListerExpansion
}

// etcdClusterLister implements the EtcdClusterLister interface.
type etcdClusterLister struct {
	indexer cache.Indexer
}

// NewEtcdClusterLister returns a new EtcdClusterLister.
func NewEtcdClusterLister(indexer cache.Indexer) EtcdClusterLister {
	return &etcdClusterLister{indexer: indexer}
}

// List lists all EtcdClusters in the indexer.
func (s *etcdClusterLister) List(selector labels.Selector) (ret []*v1alpha2.EtcdCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.EtcdCluster))
	})
	return ret, err
}

// EtcdClusters returns an object that can list and get EtcdClusters.
func (s *etcdClusterLister) EtcdClusters(namespace string) EtcdClusterNamespaceLister {
	return etcdClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EtcdClusterNamespaceLister helps list and get EtcdClusters.
// All objects returned here must be treated as read-only.
type EtcdClusterNamespaceLister interface {
	// List lists all EtcdClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.EtcdCluster, err error)
	// Get retrieves the EtcdCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.EtcdCluster, error)
	EtcdClusterNamespaceListerExpansion
}

// etcdClusterNamespaceLister implements the EtcdClusterNamespaceLister
// interface.
type etcdClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EtcdClusters in the indexer for a given namespace.
func (s etcdClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.EtcdCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.EtcdCluster))
	})
	return ret, err
}

// Get retrieves the EtcdCluster from the indexer for a given namespace and name.
func (s etcdClusterNamespaceLister) Get(name string) (*v1alpha2.EtcdCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("etcdcluster"), name)
	}
	return obj.(*v1alpha2.EtcdCluster), nil
}
