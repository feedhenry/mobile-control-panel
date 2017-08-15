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

package fake

import (
	mobile "github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMobileApps implements MobileAppInterface
type FakeMobileApps struct {
	Fake *FakeMobile
	ns   string
}

var mobileappsResource = schema.GroupVersionResource{Group: "mobile.k8s.io", Version: "", Resource: "mobileapps"}

func (c *FakeMobileApps) Create(mobileApp *mobile.MobileApp) (result *mobile.MobileApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mobileappsResource, c.ns, mobileApp), &mobile.MobileApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.MobileApp), err
}

func (c *FakeMobileApps) Update(mobileApp *mobile.MobileApp) (result *mobile.MobileApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mobileappsResource, c.ns, mobileApp), &mobile.MobileApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.MobileApp), err
}

func (c *FakeMobileApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mobileappsResource, c.ns, name), &mobile.MobileApp{})

	return err
}

func (c *FakeMobileApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mobileappsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mobile.MobileAppList{})
	return err
}

func (c *FakeMobileApps) Get(name string, options v1.GetOptions) (result *mobile.MobileApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mobileappsResource, c.ns, name), &mobile.MobileApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.MobileApp), err
}

func (c *FakeMobileApps) List(opts v1.ListOptions) (result *mobile.MobileAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mobileappsResource, c.ns, opts), &mobile.MobileAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mobile.MobileAppList{}
	for _, item := range obj.(*mobile.MobileAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mobileApps.
func (c *FakeMobileApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mobileappsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched mobileApp.
func (c *FakeMobileApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.MobileApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mobileappsResource, c.ns, name, data, subresources...), &mobile.MobileApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.MobileApp), err
}
