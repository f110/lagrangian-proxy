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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "go.f110.dev/heimdallr/operator/pkg/api/proxy/v1alpha2"
	scheme "go.f110.dev/heimdallr/operator/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackendsGetter has a method to return a BackendInterface.
// A group's client should implement this interface.
type BackendsGetter interface {
	Backends(namespace string) BackendInterface
}

// BackendInterface has methods to work with Backend resources.
type BackendInterface interface {
	Create(ctx context.Context, backend *v1alpha2.Backend, opts v1.CreateOptions) (*v1alpha2.Backend, error)
	Update(ctx context.Context, backend *v1alpha2.Backend, opts v1.UpdateOptions) (*v1alpha2.Backend, error)
	UpdateStatus(ctx context.Context, backend *v1alpha2.Backend, opts v1.UpdateOptions) (*v1alpha2.Backend, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.Backend, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.BackendList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Backend, err error)
	BackendExpansion
}

// backends implements BackendInterface
type backends struct {
	client rest.Interface
	ns     string
}

// newBackends returns a Backends
func newBackends(c *ProxyV1alpha2Client, namespace string) *backends {
	return &backends{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backend, and returns the corresponding backend object, and an error if there is any.
func (c *backends) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Backend, err error) {
	result = &v1alpha2.Backend{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backends").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Backends that match those selectors.
func (c *backends) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.BackendList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.BackendList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backends.
func (c *backends) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backend and creates it.  Returns the server's representation of the backend, and an error, if there is any.
func (c *backends) Create(ctx context.Context, backend *v1alpha2.Backend, opts v1.CreateOptions) (result *v1alpha2.Backend, err error) {
	result = &v1alpha2.Backend{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backend).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backend and updates it. Returns the server's representation of the backend, and an error, if there is any.
func (c *backends) Update(ctx context.Context, backend *v1alpha2.Backend, opts v1.UpdateOptions) (result *v1alpha2.Backend, err error) {
	result = &v1alpha2.Backend{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backends").
		Name(backend.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backend).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backends) UpdateStatus(ctx context.Context, backend *v1alpha2.Backend, opts v1.UpdateOptions) (result *v1alpha2.Backend, err error) {
	result = &v1alpha2.Backend{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backends").
		Name(backend.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backend).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backend and deletes it. Returns an error if one occurs.
func (c *backends) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backends").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backends) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backend.
func (c *backends) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Backend, err error) {
	result = &v1alpha2.Backend{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backends").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
