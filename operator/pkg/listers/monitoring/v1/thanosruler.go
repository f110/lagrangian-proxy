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

package v1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ThanosRulerLister helps list ThanosRulers.
// All objects returned here must be treated as read-only.
type ThanosRulerLister interface {
	// List lists all ThanosRulers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ThanosRuler, err error)
	// ThanosRulers returns an object that can list and get ThanosRulers.
	ThanosRulers(namespace string) ThanosRulerNamespaceLister
	ThanosRulerListerExpansion
}

// thanosRulerLister implements the ThanosRulerLister interface.
type thanosRulerLister struct {
	indexer cache.Indexer
}

// NewThanosRulerLister returns a new ThanosRulerLister.
func NewThanosRulerLister(indexer cache.Indexer) ThanosRulerLister {
	return &thanosRulerLister{indexer: indexer}
}

// List lists all ThanosRulers in the indexer.
func (s *thanosRulerLister) List(selector labels.Selector) (ret []*v1.ThanosRuler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ThanosRuler))
	})
	return ret, err
}

// ThanosRulers returns an object that can list and get ThanosRulers.
func (s *thanosRulerLister) ThanosRulers(namespace string) ThanosRulerNamespaceLister {
	return thanosRulerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ThanosRulerNamespaceLister helps list and get ThanosRulers.
// All objects returned here must be treated as read-only.
type ThanosRulerNamespaceLister interface {
	// List lists all ThanosRulers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ThanosRuler, err error)
	// Get retrieves the ThanosRuler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ThanosRuler, error)
	ThanosRulerNamespaceListerExpansion
}

// thanosRulerNamespaceLister implements the ThanosRulerNamespaceLister
// interface.
type thanosRulerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ThanosRulers in the indexer for a given namespace.
func (s thanosRulerNamespaceLister) List(selector labels.Selector) (ret []*v1.ThanosRuler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ThanosRuler))
	})
	return ret, err
}

// Get retrieves the ThanosRuler from the indexer for a given namespace and name.
func (s thanosRulerNamespaceLister) Get(name string) (*v1.ThanosRuler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("thanosruler"), name)
	}
	return obj.(*v1.ThanosRuler), nil
}
