package operations

import (
	"net/http"

	openservicebroker "github.com/feedhenry/mobile-control-panel/server/pkg/mobile/broker"
)

// Catalog is an implementation of the service broker catalog endpoint.
// This implementation returns a static set of services and plans.
func (b *BrokerOperations) Catalog() *openservicebroker.Response {
	services := make([]*openservicebroker.Service, 2)
	services[0] = &openservicebroker.Service{
		Name:        "android-app",
		ID:          "serviceUUID",
		Description: "start an android application.",
		Tags:        []string{"mobile", "android"},
		Bindable:    false,
		Metadata: map[string]interface{}{
			"displayName":                    "Android App",
			"console.openshift.io/iconClass": "fa fa-android",
		},
		Plans: []openservicebroker.Plan{
			{
				Name:        "gold-plan",
				ID:          "gold-plan-id",
				Description: "gold plan description",
				Bindable:    false,
				Free:        true,
				Schema: &openservicebroker.Schema{
					ServiceInstances: &openservicebroker.ServiceInstanceSchema{
						Create: &openservicebroker.CreateSchema{
							Parameters: &openservicebroker.Create{
								Schema:              "http://json-schema.org/draft-04/schema",
								AddtionalProperties: false,
								Properties: map[string]map[string]string{
									"name": {
										"title": "The name of your app",
										"type":  "string",
									},
									"type": {
										"default": "android",
										"title":   "The app type",
										"type":    "string",
									},
								},
								Required: []string{"name", "type"},
								Type:     "object",
							},
						},
					},
				},
			},
		},
	}
	services[1] = &openservicebroker.Service{
		Name:        "ios-app",
		ID:          "serviceIodUUID",
		Description: "start an ios application.",
		Tags:        []string{"mobile", "ios", "apple"},
		Bindable:    false,
		Metadata: map[string]interface{}{
			"displayName":                    "iOS App",
			"console.openshift.io/iconClass": "fa fa-apple",
		},
		Plans: []openservicebroker.Plan{
			{
				Name:        "gold-plan",
				ID:          "gold-plan-id",
				Description: "gold plan description",
				Bindable:    false,
				Free:        true,
				Schema: &openservicebroker.Schema{
					ServiceInstances: &openservicebroker.ServiceInstanceSchema{
						Create: &openservicebroker.CreateSchema{
							Parameters: &openservicebroker.Create{
								Schema:              "http://json-schema.org/draft-04/schema",
								AddtionalProperties: false,
								Properties: map[string]map[string]string{
									"name": {
										"title": "The name of your app",
										"type":  "string",
									},
									"type": {
										"default": "ios",
										"title":   "The app type",
										"type":    "string",
									},
								},
								Required: []string{"name", "type"},
								Type:     "object",
							},
						},
					},
				},
			},
		},
	}
	return &openservicebroker.Response{Code: http.StatusOK, Body: &openservicebroker.CatalogResponse{Services: services}, Err: nil}
}
