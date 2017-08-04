package operations

import (
	"net/http"

	openservicebroker "github.com/feedhenry/mobile-control-panel/server/pkg/mobile/broker"
)

// Catalog is an implementation of the service broker catalog endpoint.
// This implementation returns a static set of services and plans.
func (b *BrokerOperations) Catalog() *openservicebroker.Response {

	services := make([]*openservicebroker.Service, 2)

	/**

				alphaInstanceCreateParameterSchema": {
			            "$schema": "http://json-schema.org/draft-04/schema",
			            "additionalProperties": false,
			            "properties": {
			              "ENABLE_OAUTH": {
			                "default": "true",
			                "description": "Whether to enable OAuth OpenShift integration. If false, the static account 'admin' will be initialized with the password 'password'.",
			                "title": "Enable OAuth in Jenkins",
			                "type": "string"
						  },
									   "required": [
		              "template.openshift.io/requester-username"
		            ],
					"type": "object"
								"default": "true",
	                "description": "Whether to enable OAuth OpenShift integration. If false, the static account 'admin' will be initialized with the password 'password'.",
	                "title": "Enable OAuth in Jenkins",
	                "type": "string"
				**/

	services[0] = &openservicebroker.Service{
		Name:        "android-app",
		ID:          "serviceUUID",
		Description: "start an android application.",
		Tags:        []string{"mobile", "android"},
		Bindable:    false,
		Metadata: map[string]interface{}{
			"displayName": "Android Starter",
		},
		Plans: []openservicebroker.Plan{
			openservicebroker.Plan{
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
									"name": map[string]string{
										"title": "The name of your app",
										"type":  "string",
									},
									"type": map[string]string{
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
		Tags:        []string{"mobile", "android"},
		Bindable:    false,
		Metadata: map[string]interface{}{
			"displayName": "iOS Starter",
		},
		Plans: []openservicebroker.Plan{
			openservicebroker.Plan{
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
									"name": map[string]string{
										"title": "The name of your app",
										"type":  "string",
									},
									"type": map[string]string{
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
