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

package fake

import (
	"context"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceMonitors implements ServiceMonitorInterface
type FakeServiceMonitors struct {
	Fake *FakeMonitoringV1
	ns   string
}

var servicemonitorsResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "servicemonitors"}

var servicemonitorsKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "ServiceMonitor"}

// Get takes name of the serviceMonitor, and returns the corresponding serviceMonitor object, and an error if there is any.
func (c *FakeServiceMonitors) Get(ctx context.Context, name string, options v1.GetOptions) (result *monitoringv1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemonitorsResource, c.ns, name), &monitoringv1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ServiceMonitor), err
}

// List takes label and field selectors, and returns the list of ServiceMonitors that match those selectors.
func (c *FakeServiceMonitors) List(ctx context.Context, opts v1.ListOptions) (result *monitoringv1.ServiceMonitorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemonitorsResource, servicemonitorsKind, c.ns, opts), &monitoringv1.ServiceMonitorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoringv1.ServiceMonitorList{ListMeta: obj.(*monitoringv1.ServiceMonitorList).ListMeta}
	for _, item := range obj.(*monitoringv1.ServiceMonitorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceMonitors.
func (c *FakeServiceMonitors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemonitorsResource, c.ns, opts))

}

// Create takes the representation of a serviceMonitor and creates it.  Returns the server's representation of the serviceMonitor, and an error, if there is any.
func (c *FakeServiceMonitors) Create(ctx context.Context, serviceMonitor *monitoringv1.ServiceMonitor, opts v1.CreateOptions) (result *monitoringv1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemonitorsResource, c.ns, serviceMonitor), &monitoringv1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ServiceMonitor), err
}

// Update takes the representation of a serviceMonitor and updates it. Returns the server's representation of the serviceMonitor, and an error, if there is any.
func (c *FakeServiceMonitors) Update(ctx context.Context, serviceMonitor *monitoringv1.ServiceMonitor, opts v1.UpdateOptions) (result *monitoringv1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemonitorsResource, c.ns, serviceMonitor), &monitoringv1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ServiceMonitor), err
}

// Delete takes name of the serviceMonitor and deletes it. Returns an error if one occurs.
func (c *FakeServiceMonitors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicemonitorsResource, c.ns, name), &monitoringv1.ServiceMonitor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMonitors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemonitorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &monitoringv1.ServiceMonitorList{})
	return err
}

// Patch applies the patch and returns the patched serviceMonitor.
func (c *FakeServiceMonitors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *monitoringv1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemonitorsResource, c.ns, name, pt, data, subresources...), &monitoringv1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ServiceMonitor), err
}
