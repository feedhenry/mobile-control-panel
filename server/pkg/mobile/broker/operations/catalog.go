package operations

import (
	"net/http"

	openservicebroker "github.com/feedhenry/mobile-control-panel/server/pkg/mobile/broker"
)

// Catalog is an implementation of the service broker catalog endpoint.
// This implementation returns a static set of services and plans.
func (b *BrokerOperations) Catalog() *openservicebroker.Response {

	services := make([]*openservicebroker.Service, 1)

	serviceMetadata := make(map[string]interface{})
	serviceMetadata["metadata_key1"] = "metadata_value1"

	servicePlans := make([]openservicebroker.Plan, 1)
	servicePlans[0] = openservicebroker.Plan{
		Name:        "gold-plan",
		ID:          "gold-plan-id",
		Description: "gold plan description",
		Bindable:    false,
		Free:        true,
	}
	services[0] = &openservicebroker.Service{
		Name:        "android-app",
		ID:          "serviceUUID",
		Description: "start an android application.",
		Tags:        []string{"tag1", "tag2"},
		Bindable:    false,
		Metadata:    serviceMetadata,
		Plans:       servicePlans,
	}
	return &openservicebroker.Response{Code: http.StatusOK, Body: &openservicebroker.CatalogResponse{Services: services}, Err: nil}
}
