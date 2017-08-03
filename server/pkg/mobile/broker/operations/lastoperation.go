package operations

import (
	"net/http"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"
	openservicebroker "github.com/feedhenry/mobile-control-panel/server/pkg/mobile/broker"
)

// LastOperation is an implementation of the service broker last operation api.
func (b *BrokerOperations) LastOperation(instanceID string, operation openservicebroker.Operation) *openservicebroker.Response {
	// Find the ServiceInstance that represents the state of this service instanceid
	_, err := b.Client.Mobile().MobileApps(mobile.Namespace).Get(instanceID, metav1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			if operation == openservicebroker.OperationDeprovisioning {
				return &openservicebroker.Response{Code: http.StatusGone, Body: &struct{}{}, Err: nil}
			}
			return &openservicebroker.Response{Code: http.StatusBadRequest, Body: nil, Err: err}

		}
		return &openservicebroker.Response{Code: http.StatusInternalServerError, Body: nil, Err: err}
	}

	// Check the conditions on the ServiceInstance to determine the operation state.
	// If there are no conditions, the controller has not processes the instance yet,
	// so it's in progress.  Otherwise there will be a ready or failed condition present.

	state := openservicebroker.LastOperationStateSucceeded

	return &openservicebroker.Response{Code: http.StatusOK, Body: &openservicebroker.LastOperationResponse{State: state}, Err: nil}
}
