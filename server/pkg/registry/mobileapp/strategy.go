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

package serviceinstance

// this was copied from where else and edited to fit our objects

import (
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"

	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	mobileapi "github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"
	"github.com/golang/glog"
)

type apiServerStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func NewStrategy(typer runtime.ObjectTyper) apiServerStrategy {
	return apiServerStrategy{typer, names.SimpleNameGenerator}
}

func (apiServerStrategy) NamespaceScoped() bool {
	return true
}

func (apiServerStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {

}

func (apiServerStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (apiServerStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	var errs = field.ErrorList{}
	app, ok := obj.(*mobileapi.MobileApp)
	if !ok {
		glog.Fatal("received a non MobileApp object to create")
	}
	ns := genericapirequest.NamespaceValue(ctx)
	fmt.Println("strategy Validate", ns)
	validAppTypes := map[string]bool{"android": true, "ios": true, "cordova": true}
	if _, ok := validAppTypes[app.Spec.ClientType]; !ok {
		errs = append(errs, field.Invalid(field.NewPath("Spec.ClientType"), app.Spec.ClientType, "invalid client type. valid types android or ios"))
	}
	return errs
}

func (apiServerStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (apiServerStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (apiServerStrategy) Canonicalize(obj runtime.Object) {
}

func (apiServerStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
	// return validation.ValidateServiceInstanceUpdate(obj.(*mobileapi.ServiceInstance), old.(*mobileapi.ServiceInstance))
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*mobileapi.MobileApp)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a MobileApp")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), MobileAppToSelectableFields(apiserver), nil
}

// MatchMobileApp is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchMobileApp(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// MobileAppToSelectableFields returns a field set that represents the object.
func MobileAppToSelectableFields(obj *mobileapi.MobileApp) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}
