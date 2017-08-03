package operations

import (
	"errors"
	"net/http"

	"github.com/golang/glog"

	"github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"
	openservicebroker "github.com/feedhenry/mobile-control-panel/server/pkg/mobile/broker"
)

// Provision is an implementation of the service broker provision api
func (b *BrokerOperations) Provision(instanceID string, preq *openservicebroker.ProvisionRequest) *openservicebroker.Response {
	if !preq.AcceptsIncomplete {
		return &openservicebroker.Response{Code: http.StatusUnprocessableEntity, Body: openservicebroker.ProvisionResponse{}, Err: errors.New("cannot provision service synchronously")}
	}
	glog.Info("Provision called for broker", instanceID)
	// provision will create a new ServiceInstance resource to be processed
	// by the controller.
	si := mobile.MobileApp{}
	si.Name = instanceID
	// this credential will be returned to bind requests, in theory it is a value
	// consumers of the service instance will need to access the instance.

	//si.Status.Conditions = append(si.Status.Conditions, brokerapi.ServiceInstanceCondition{})

	// Create the ServiceInstance object that represents this service instance.  The
	// controller will see the request and progress it from there.
	glog.Info("creating mobile app", preq.Parameters, preq.SpaceID)
	_, err := b.Client.Mobile().MobileApps(mobile.Namespace).Create(&si)
	if err != nil {
		glog.Errorf("Failed to create a service instance\n:%v\n", err)
		return &openservicebroker.Response{Code: http.StatusInternalServerError, Body: nil, Err: err}
	}

	// Use this for async provision flows.  Technically the service instance isn't provisioned
	// until the controller sees the request, does work, and marks it ready.
	return &openservicebroker.Response{Code: http.StatusAccepted, Body: openservicebroker.ProvisionResponse{Operation: openservicebroker.OperationProvisioning}, Err: err}

	// For synchronous flows we can just return complete.
	//return &openservicebroker.Response{http.StatusOK, openservicebroker.ProvisionResponse{Operation: openservicebroker.OperationProvisioning}, nil}
}
