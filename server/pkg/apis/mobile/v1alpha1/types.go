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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// MobileAppList is a list of MobileApp objects.
type MobileAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []MobileApp `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient=true

// MobileApp represents a service instance provision request,
// possibly fullfilled.
type MobileApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec MobileAppSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// MobileAppSpec defines the requested MobileApp
type MobileAppSpec struct {
	// Credential is a sample value associatd w/ the provisioned service
	ClientType string `json:"clientType" protobuf:"bytes,1,opt,name=clientType"`
}

const (
	APIGroupVersion = "v1alpha1"
)
