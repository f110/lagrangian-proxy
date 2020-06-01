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
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	scheme "github.com/f110/lagrangian-proxy/operator/pkg/client/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PrometheusRulesGetter has a method to return a PrometheusRuleInterface.
// A group's client should implement this interface.
type PrometheusRulesGetter interface {
	PrometheusRules(namespace string) PrometheusRuleInterface
}

// PrometheusRuleInterface has methods to work with PrometheusRule resources.
type PrometheusRuleInterface interface {
	Create(*v1.PrometheusRule) (*v1.PrometheusRule, error)
	Update(*v1.PrometheusRule) (*v1.PrometheusRule, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.PrometheusRule, error)
	List(opts metav1.ListOptions) (*v1.PrometheusRuleList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PrometheusRule, err error)
	PrometheusRuleExpansion
}

// prometheusRules implements PrometheusRuleInterface
type prometheusRules struct {
	client rest.Interface
	ns     string
}

// newPrometheusRules returns a PrometheusRules
func newPrometheusRules(c *MonitoringV1Client, namespace string) *prometheusRules {
	return &prometheusRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the prometheusRule, and returns the corresponding prometheusRule object, and an error if there is any.
func (c *prometheusRules) Get(name string, options metav1.GetOptions) (result *v1.PrometheusRule, err error) {
	result = &v1.PrometheusRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheusrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PrometheusRules that match those selectors.
func (c *prometheusRules) List(opts metav1.ListOptions) (result *v1.PrometheusRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PrometheusRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheusrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested prometheusRules.
func (c *prometheusRules) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("prometheusrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a prometheusRule and creates it.  Returns the server's representation of the prometheusRule, and an error, if there is any.
func (c *prometheusRules) Create(prometheusRule *v1.PrometheusRule) (result *v1.PrometheusRule, err error) {
	result = &v1.PrometheusRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("prometheusrules").
		Body(prometheusRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a prometheusRule and updates it. Returns the server's representation of the prometheusRule, and an error, if there is any.
func (c *prometheusRules) Update(prometheusRule *v1.PrometheusRule) (result *v1.PrometheusRule, err error) {
	result = &v1.PrometheusRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prometheusrules").
		Name(prometheusRule.Name).
		Body(prometheusRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the prometheusRule and deletes it. Returns an error if one occurs.
func (c *prometheusRules) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheusrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *prometheusRules) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheusrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched prometheusRule.
func (c *prometheusRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PrometheusRule, err error) {
	result = &v1.PrometheusRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("prometheusrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
