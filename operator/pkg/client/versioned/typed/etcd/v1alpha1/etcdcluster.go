/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/f110/lagrangian-proxy/operator/pkg/api/etcd/v1alpha1"
	scheme "github.com/f110/lagrangian-proxy/operator/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EtcdClustersGetter has a method to return a EtcdClusterInterface.
// A group's client should implement this interface.
type EtcdClustersGetter interface {
	EtcdClusters(namespace string) EtcdClusterInterface
}

// EtcdClusterInterface has methods to work with EtcdCluster resources.
type EtcdClusterInterface interface {
	Create(*v1alpha1.EtcdCluster) (*v1alpha1.EtcdCluster, error)
	Update(*v1alpha1.EtcdCluster) (*v1alpha1.EtcdCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EtcdCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.EtcdClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EtcdCluster, err error)
	EtcdClusterExpansion
}

// etcdClusters implements EtcdClusterInterface
type etcdClusters struct {
	client rest.Interface
	ns     string
}

// newEtcdClusters returns a EtcdClusters
func newEtcdClusters(c *EtcdV1alpha1Client, namespace string) *etcdClusters {
	return &etcdClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the etcdCluster, and returns the corresponding etcdCluster object, and an error if there is any.
func (c *etcdClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.EtcdCluster, err error) {
	result = &v1alpha1.EtcdCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("etcdclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EtcdClusters that match those selectors.
func (c *etcdClusters) List(opts v1.ListOptions) (result *v1alpha1.EtcdClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EtcdClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("etcdclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested etcdClusters.
func (c *etcdClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("etcdclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a etcdCluster and creates it.  Returns the server's representation of the etcdCluster, and an error, if there is any.
func (c *etcdClusters) Create(etcdCluster *v1alpha1.EtcdCluster) (result *v1alpha1.EtcdCluster, err error) {
	result = &v1alpha1.EtcdCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("etcdclusters").
		Body(etcdCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a etcdCluster and updates it. Returns the server's representation of the etcdCluster, and an error, if there is any.
func (c *etcdClusters) Update(etcdCluster *v1alpha1.EtcdCluster) (result *v1alpha1.EtcdCluster, err error) {
	result = &v1alpha1.EtcdCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("etcdclusters").
		Name(etcdCluster.Name).
		Body(etcdCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the etcdCluster and deletes it. Returns an error if one occurs.
func (c *etcdClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("etcdclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *etcdClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("etcdclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched etcdCluster.
func (c *etcdClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EtcdCluster, err error) {
	result = &v1alpha1.EtcdCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("etcdclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
