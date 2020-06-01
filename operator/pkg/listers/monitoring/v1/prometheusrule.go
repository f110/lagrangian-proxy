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
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PrometheusRuleLister helps list PrometheusRules.
type PrometheusRuleLister interface {
	// List lists all PrometheusRules in the indexer.
	List(selector labels.Selector) (ret []*v1.PrometheusRule, err error)
	// PrometheusRules returns an object that can list and get PrometheusRules.
	PrometheusRules(namespace string) PrometheusRuleNamespaceLister
	PrometheusRuleListerExpansion
}

// prometheusRuleLister implements the PrometheusRuleLister interface.
type prometheusRuleLister struct {
	indexer cache.Indexer
}

// NewPrometheusRuleLister returns a new PrometheusRuleLister.
func NewPrometheusRuleLister(indexer cache.Indexer) PrometheusRuleLister {
	return &prometheusRuleLister{indexer: indexer}
}

// List lists all PrometheusRules in the indexer.
func (s *prometheusRuleLister) List(selector labels.Selector) (ret []*v1.PrometheusRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PrometheusRule))
	})
	return ret, err
}

// PrometheusRules returns an object that can list and get PrometheusRules.
func (s *prometheusRuleLister) PrometheusRules(namespace string) PrometheusRuleNamespaceLister {
	return prometheusRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrometheusRuleNamespaceLister helps list and get PrometheusRules.
type PrometheusRuleNamespaceLister interface {
	// List lists all PrometheusRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.PrometheusRule, err error)
	// Get retrieves the PrometheusRule from the indexer for a given namespace and name.
	Get(name string) (*v1.PrometheusRule, error)
	PrometheusRuleNamespaceListerExpansion
}

// prometheusRuleNamespaceLister implements the PrometheusRuleNamespaceLister
// interface.
type prometheusRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrometheusRules in the indexer for a given namespace.
func (s prometheusRuleNamespaceLister) List(selector labels.Selector) (ret []*v1.PrometheusRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PrometheusRule))
	})
	return ret, err
}

// Get retrieves the PrometheusRule from the indexer for a given namespace and name.
func (s prometheusRuleNamespaceLister) Get(name string) (*v1.PrometheusRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("prometheusrule"), name)
	}
	return obj.(*v1.PrometheusRule), nil
}
