/*
Copyright 2021 The KubeVirt Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "kubevirt.io/client-go/apis/snapshot/v1alpha1"
	scheme "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/scheme"
)

// VirtualMachineSnapshotContentsGetter has a method to return a VirtualMachineSnapshotContentInterface.
// A group's client should implement this interface.
type VirtualMachineSnapshotContentsGetter interface {
	VirtualMachineSnapshotContents(namespace string) VirtualMachineSnapshotContentInterface
}

// VirtualMachineSnapshotContentInterface has methods to work with VirtualMachineSnapshotContent resources.
type VirtualMachineSnapshotContentInterface interface {
	Create(*v1alpha1.VirtualMachineSnapshotContent) (*v1alpha1.VirtualMachineSnapshotContent, error)
	Update(*v1alpha1.VirtualMachineSnapshotContent) (*v1alpha1.VirtualMachineSnapshotContent, error)
	UpdateStatus(*v1alpha1.VirtualMachineSnapshotContent) (*v1alpha1.VirtualMachineSnapshotContent, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualMachineSnapshotContent, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualMachineSnapshotContentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineSnapshotContent, err error)
	VirtualMachineSnapshotContentExpansion
}

// virtualMachineSnapshotContents implements VirtualMachineSnapshotContentInterface
type virtualMachineSnapshotContents struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineSnapshotContents returns a VirtualMachineSnapshotContents
func newVirtualMachineSnapshotContents(c *SnapshotV1alpha1Client, namespace string) *virtualMachineSnapshotContents {
	return &virtualMachineSnapshotContents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineSnapshotContent, and returns the corresponding virtualMachineSnapshotContent object, and an error if there is any.
func (c *virtualMachineSnapshotContents) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineSnapshotContents that match those selectors.
func (c *virtualMachineSnapshotContents) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineSnapshotContentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineSnapshotContentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineSnapshotContents.
func (c *virtualMachineSnapshotContents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineSnapshotContent and creates it.  Returns the server's representation of the virtualMachineSnapshotContent, and an error, if there is any.
func (c *virtualMachineSnapshotContents) Create(virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Body(virtualMachineSnapshotContent).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineSnapshotContent and updates it. Returns the server's representation of the virtualMachineSnapshotContent, and an error, if there is any.
func (c *virtualMachineSnapshotContents) Update(virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(virtualMachineSnapshotContent.Name).
		Body(virtualMachineSnapshotContent).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineSnapshotContents) UpdateStatus(virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(virtualMachineSnapshotContent.Name).
		SubResource("status").
		Body(virtualMachineSnapshotContent).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineSnapshotContent and deletes it. Returns an error if one occurs.
func (c *virtualMachineSnapshotContents) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineSnapshotContents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineSnapshotContent.
func (c *virtualMachineSnapshotContents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
