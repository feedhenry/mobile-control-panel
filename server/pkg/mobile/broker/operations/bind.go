package operations

import (
	"net/http"

	openservicebroker "github.com/feedhenry/mobile-control-panel/server/pkg/mobile/broker"
)

// Bind handles bind requests from the service catalog by returning
// a bind response with credentials for the service instance.
func (b *BrokerOperations) Bind(instanceID, bindingID string, breq *openservicebroker.BindRequest) *openservicebroker.Response {
	// Find the service instance that is being bound to

	// in principle, bind should alter state somewhere

	// Create some credentials to return.  In this case the credentials are
	// pulled from the service instance but a real broker might
	// return unique credentials for each binding so that multiple users
	// of a service instance are not sharing credentials.
	credentials := map[string]interface{}{}
	//credentials["credential"] = si.Spec.Credential

	return &openservicebroker.Response{
		Code: http.StatusCreated,
		Body: &openservicebroker.BindResponse{Credentials: credentials},
		Err:  nil,
	}
}
