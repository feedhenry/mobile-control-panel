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

package controller

import (
	"github.com/golang/glog"

	mobileapi "github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"
	mobileclientset "github.com/feedhenry/mobile-control-panel/server/pkg/client/clientset_generated/internalclientset"
	"github.com/feedhenry/mobile-control-panel/server/pkg/client/informers_generated/internalversion/mobile/internalversion"
	"k8s.io/client-go/tools/cache"
)

// Controller describes a controller that processes service instance
// provision requests for the broker.
type Controller interface {
	// Run runs the controller until the given stop channel can be read from.
	Run(stopCh <-chan struct{})
}

// controller is a concrete Controller.
type controller struct {
	mobileClient mobileclientset.Clientset
	informer     cache.SharedIndexInformer
}

// New returns a new Open Service Broker provision controller.
func New(mobileClient mobileclientset.Clientset, informers internalversion.MobileAppInformer) (Controller, error) {

	controller := &controller{
		mobileClient: mobileClient,
		informer:     informers.Informer(),
	}

	// setup an informer that will tell us about new/updated/deleted MobileApp objects.
	// (we don't actually do anything w/ updated objects in this controller)
	controller.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			glog.Info("add func called ", obj)
		},
		UpdateFunc: func(_, obj interface{}) {

		},
		DeleteFunc: func(obj interface{}) {
		}})

	return controller, nil
}

// Run runs the controller until the given stop channel can be read from.
func (c *controller) Run(stopCh <-chan struct{}) {
	glog.Info("Starting mobile controller")
	c.informer.Run(stopCh)
}

// serviceInstanceAdd handles added ServiceInstances.
// It will update the service instance to indicate it has
// been successfully provisioned/is ready for use.
// A real broker could use this to create any backing
// service resources needed to run the service instance.
func (c *controller) mobileAppAdd(obj interface{}) {
	instance, ok := obj.(*mobileapi.MobileApp)
	if instance == nil || !ok {
		return
	}

	// Controllers periodically get a full relist of all resources they are watching,
	// so we need to make sure we properly no-op service instances we've already
	// handled previously.
	// for _, condition := range instance.Status.Conditions {
	// 	if condition.Type == mobileapi.MobileAppReady && condition.Status == kapi.ConditionTrue {
	// 		// This provision request has already been fulfilled.
	// 		return
	// 	}
	// 	if condition.Type == mobileapi.MobileAppFailed && condition.Status == kapi.ConditionTrue {
	// 		// This provision request has already failed.
	// 		return
	// 	}
	// }

	// glog.Infof("controller processing provision request for instance %s", instance.Name)
	// // add the Ready condition to the service instance.
	// condition := mobileapi.MobileAppCondition{
	// 	Type:               mobileapi.MobileAppReady,
	// 	Status:             kapi.ConditionTrue,
	// 	LastTransitionTime: metav1.Now(),
	// 	Reason:             "ServiceProvisioned",
	// 	Message:            "This service has been provisioned",
	// }
	// instance.Status.Conditions = append(instance.Status.Conditions, condition)
	// _, err := c.brokerClient.Broker().ServiceInstances("brokersdk").Update(instance)
	// if err != nil {
	// 	glog.Errorf("Error updating service instance %s to ready: %v", instance.Name, err)
	// }
}

// serviceInstanceDelete handles deleted ServiceInstances
// Currently there is no action that needs to be taken when
// a service instance is deleted, but a real controller could
// use this to tear down the backing service resources.
func (c *controller) serviceInstanceDelete(obj interface{}) {
	instance, ok := obj.(*mobileapi.MobileApp)
	if instance == nil || !ok {
		return
	}
	glog.Infof("controller processing deprovision request for instance %s", instance.Name)

}
