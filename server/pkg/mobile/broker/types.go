package broker

// from https://github.com/openservicebrokerapi/servicebroker/blob/1d301105c66187b5aa2e061a1264ecf3cbc3d2a0/_spec.md

/* These types represent the open service broker api spec */

const (
	// XBrokerAPIVersion is the header for the openservice broker api versoin
	XBrokerAPIVersion = "X-Broker-Api-Version"
	// APIVersion is the version of the openservice broker api supported by this broker
	APIVersion      = "2.11"
	BrokerAPIPrefix = "/broker/mobile.broker.io"
)

// Service is an available service listed in the catalog
type Service struct {
	Name            string                 `json:"name"`
	ID              string                 `json:"id"`
	Description     string                 `json:"description"`
	Tags            []string               `json:"tags,omitempty"`
	Requires        []string               `json:"requires,omitempty"`
	Bindable        bool                   `json:"bindable"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	DashboardClient *DashboardClient       `json:"dashboard_client,omitempty"`
	PlanUpdatable   bool                   `json:"plan_updateable,omitempty"`
	Plans           []Plan                 `json:"plans"`
}

// DashboardClient contains information about the oauth SSO
// flow for the service dashboard
type DashboardClient struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	RedirectURI string `json:"redirect_uri"`
}

// Create this is custom to make it work with the Mall should be removed if we do not need it in the future.
type Create struct {
	Schema              string                       `json:"$schema"`
	AddtionalProperties bool                         `json:"additionalProperties"`
	Properties          map[string]map[string]string `json:"properties"`
	Required            []string                     `json:"required"`
	Type                string                       `json:"type"`
}

type CreateSchema struct {
	Parameters interface{} `json:"parameters,omitempty"`
}

// Schema represents a plan's schemas for service instance and binding create
// and update.
type Schema struct {
	ServiceInstances *ServiceInstanceSchema `json:"service_instance,omitempty"`
}

// ServiceInstanceSchema represents a plan's schemas for a create and update
// of a service instance.
type ServiceInstanceSchema struct {
	Create *CreateSchema `json:"create,omitempty"`
	Update *CreateSchema `json:"update,omitempty"`
}

// Plan is a plan within a service offering
type Plan struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Free        bool                   `json:"free,omitempty"`
	Bindable    bool                   `json:"bindable,omitempty"`
	Schema      *Schema                `json:"schemas,omitempty"`
}

// CatalogResponse is sent as the response to catalog requests
type CatalogResponse struct {
	Services []*Service `json:"services"`
}

// LastOperationResponse is sent as a response to last operation requests
type LastOperationResponse struct {
	State       LastOperationState `json:"state"`
	Description string             `json:"description,omitempty"`
}

type LastOperationState string

const (
	LastOperationStateInProgress LastOperationState = "in progress"
	LastOperationStateSucceeded  LastOperationState = "succeeded"
	LastOperationStateFailed     LastOperationState = "failed"
)

type KubernetesContext struct {
	Platform  string `json:"platform"`
	Namespace string `json:"namespace"`
}

// ProvisionRequest is sent as part of a provision api call
type ProvisionRequest struct {
	Context           KubernetesContext `json:"context"`
	ServiceID         string            `json:"service_id"`
	PlanID            string            `json:"plan_id"`
	Parameters        map[string]string `json:"parameters,omitempty"`
	AcceptsIncomplete bool              `json:"accepts_incomplete,omitempty"`
	OrganizationID    string            `json:"organization_guid"`
	SpaceID           string            `json:"space_guid"`
}

// ProvisionResponse is sent in response to a provision call
type ProvisionResponse struct {
	DashboardURL string    `json:"dashboard_url,omitempty"`
	Operation    Operation `json:"operation,omitempty"`
}

type Operation string

type UpdateRequest struct {
	ServiceID         string            `json:"service_id"`
	PlanID            string            `json:"plan_id,omitempty"`
	Parameters        map[string]string `json:"parameters,omitempty"`
	AcceptsIncomplete bool              `json:"accepts_incomplete,omitempty"`
	PreviousValues    struct {
		ServiceID      string `json:"service_id,omitempty"`
		PlanID         string `json:"plan_id,omitempty"`
		OrganizationID string `json:"organization_id,omitempty"`
		SpaceID        string `json:"space_id,omitempty"`
	} `json:"previous_values,omitempty"`
}

type UpdateResponse struct {
	Operation Operation `json:"operation,omitempty"`
}

// BindRequest is sent as part of a bind api call
type BindRequest struct {
	ServiceID    string `json:"service_id"`
	PlanID       string `json:"plan_id"`
	AppGUID      string `json:"app_guid,omitempty"`
	BindResource struct {
		AppGUID string `json:"app_guid,omitempty"`
		Route   string `json:"route,omitempty"`
	} `json:"bind_resource,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// BindResponse is sent in response to a bind api call
type BindResponse struct {
	Credentials     map[string]interface{} `json:"credentials,omitempty"`
	SyslogDrainURL  string                 `json:"syslog_drain_url,omitempty"`
	RouteServiceURL string                 `json:"route_service_url,omitempty"`
	VolumeMounts    []interface{}          `json:"volume_mounts,omitempty"`
}

// UnbindResponse is sent in response to an unbind call
type UnbindResponse struct {
}

// DeprovisionResponse is sent in response to a deprovision call
type DeprovisionResponse struct {
	Operation Operation `json:"operation,omitempty"`
}

type ErrorResponse struct {
	Description string `json:"description"`
}

var AsyncRequired = struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"description"`
}{
	Error:       "AsyncRequired",
	Description: "This service plan requires client support for asynchronous service operations.",
}

// from http://docs.cloudfoundry.org/services/catalog-metadata.html#services-metadata-fields

const (
	ServiceMetadataDisplayName         = "displayName"
	ServiceMetadataImageURL            = "imageUrl"
	ServiceMetadataLongDescription     = "longDescription"
	ServiceMetadataProviderDisplayName = "providerDisplayName"
	ServiceMetadataDocumentationURL    = "documentationUrl"
	ServiceMetadataSupportURL          = "supportUrl"
)

// the types below are not specified in the openservicebrokerapi spec

type Response struct {
	Code int
	Body interface{}
	Err  error
}

type Broker interface {
	Catalog() *Response
	Provision(instance_id string, preq *ProvisionRequest) *Response
	Deprovision(instance_id string) *Response
	Bind(instance_id string, binding_id string, breq *BindRequest) *Response
	Unbind(instance_id string, binding_id string) *Response
	LastOperation(instance_id string, operation Operation) *Response
}

const (
	OperationProvisioning   Operation = "provisioning"
	OperationUpdating       Operation = "updating"
	OperationDeprovisioning Operation = "deprovisioning"
)
