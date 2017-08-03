/*
Copyright 2017 The Kubernetes Authors.

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

package internalversion

import (
	mobile "github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"
	scheme "github.com/feedhenry/mobile-control-panel/server/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MobileAppsGetter has a method to return a MobileAppInterface.
// A group's client should implement this interface.
type MobileAppsGetter interface {
	MobileApps(namespace string) MobileAppInterface
}

// MobileAppInterface has methods to work with MobileApp resources.
type MobileAppInterface interface {
	Create(*mobile.MobileApp) (*mobile.MobileApp, error)
	Update(*mobile.MobileApp) (*mobile.MobileApp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*mobile.MobileApp, error)
	List(opts v1.ListOptions) (*mobile.MobileAppList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.MobileApp, err error)
	MobileAppExpansion
}

// mobileApps implements MobileAppInterface
type mobileApps struct {
	client rest.Interface
	ns     string
}

// newMobileApps returns a MobileApps
func newMobileApps(c *MobileClient, namespace string) *mobileApps {
	return &mobileApps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a mobileApp and creates it.  Returns the server's representation of the mobileApp, and an error, if there is any.
func (c *mobileApps) Create(mobileApp *mobile.MobileApp) (result *mobile.MobileApp, err error) {
	result = &mobile.MobileApp{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mobileapps").
		Body(mobileApp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mobileApp and updates it. Returns the server's representation of the mobileApp, and an error, if there is any.
func (c *mobileApps) Update(mobileApp *mobile.MobileApp) (result *mobile.MobileApp, err error) {
	result = &mobile.MobileApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mobileapps").
		Name(mobileApp.Name).
		Body(mobileApp).
		Do().
		Into(result)
	return
}

// Delete takes name of the mobileApp and deletes it. Returns an error if one occurs.
func (c *mobileApps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mobileapps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mobileApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mobileapps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the mobileApp, and returns the corresponding mobileApp object, and an error if there is any.
func (c *mobileApps) Get(name string, options v1.GetOptions) (result *mobile.MobileApp, err error) {
	result = &mobile.MobileApp{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mobileapps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MobileApps that match those selectors.
func (c *mobileApps) List(opts v1.ListOptions) (result *mobile.MobileAppList, err error) {
	result = &mobile.MobileAppList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mobileapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mobileApps.
func (c *mobileApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mobileapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched mobileApp.
func (c *mobileApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.MobileApp, err error) {
	result = &mobile.MobileApp{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mobileapps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
