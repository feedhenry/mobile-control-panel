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

package mobile

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// MobileAppList is a list of MobileApp objects.
type MobileAppList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []MobileApp
}

// +genclient=true

// MobileApp represents a service instance provision request,
// possibly fullfilled.
type MobileApp struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec MobileAppSpec
}

// MobileAppSpec defines the requested MobileApp
type MobileAppSpec struct {
	ClientType string
}

const (

	// TypePackage is the name of the package that defines the resource types
	// used by this broker.
	TypePackage = "github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"

	// GroupName is the name of the api group used for resources created/managed
	// by this broker.
	GroupName = "mobile.k8s.io"

	// MobileAppsResource is the name of the resource used to represent
	// provision requests(possibly fulfilled) for service instances
	MobileAppsResource = "mobileapps"

	// MobileAppResource is the name of the resource used to represent
	// provision requests(possibly fulfilled) for service instances
	MobileAppResource = "mobileapp"

	// MobileAPIPrefix is the route prefix for the open service broker api
	// endpoints (e.g. https://yourhost.com/mobile/mobile.srv.io/v2/catalog)
	MobileAPIPrefix = "/mobile/mobile.srv.io"

	// Namespace is the namespace the broker will be deployed in and
	// under which it will create any resources
	Namespace = "mobile"
)
