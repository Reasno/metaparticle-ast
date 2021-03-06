// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/metaparticle-io/metaparticle-ast/restapi/operations/services"
)

// NewAnApplicationForEasierDistributedApplicationGenerationAPI creates a new AnApplicationForEasierDistributedApplicationGeneration instance
func NewAnApplicationForEasierDistributedApplicationGenerationAPI(spec *loads.Document) *AnApplicationForEasierDistributedApplicationGenerationAPI {
	return &AnApplicationForEasierDistributedApplicationGenerationAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ServicesCreateOrUpdateServiceHandler: services.CreateOrUpdateServiceHandlerFunc(func(params services.CreateOrUpdateServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation ServicesCreateOrUpdateService has not yet been implemented")
		}),
		ServicesDeleteServiceHandler: services.DeleteServiceHandlerFunc(func(params services.DeleteServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation ServicesDeleteService has not yet been implemented")
		}),
		ServicesGetServiceHandler: services.GetServiceHandlerFunc(func(params services.GetServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation ServicesGetService has not yet been implemented")
		}),
		ServicesListServicesHandler: services.ListServicesHandlerFunc(func(params services.ListServicesParams) middleware.Responder {
			return middleware.NotImplemented("operation ServicesListServices has not yet been implemented")
		}),
	}
}

/*AnApplicationForEasierDistributedApplicationGenerationAPI The metaparticle API */
type AnApplicationForEasierDistributedApplicationGenerationAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ServicesCreateOrUpdateServiceHandler sets the operation handler for the create or update service operation
	ServicesCreateOrUpdateServiceHandler services.CreateOrUpdateServiceHandler
	// ServicesDeleteServiceHandler sets the operation handler for the delete service operation
	ServicesDeleteServiceHandler services.DeleteServiceHandler
	// ServicesGetServiceHandler sets the operation handler for the get service operation
	ServicesGetServiceHandler services.GetServiceHandler
	// ServicesListServicesHandler sets the operation handler for the list services operation
	ServicesListServicesHandler services.ListServicesHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AnApplicationForEasierDistributedApplicationGenerationAPI
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ServicesCreateOrUpdateServiceHandler == nil {
		unregistered = append(unregistered, "services.CreateOrUpdateServiceHandler")
	}

	if o.ServicesDeleteServiceHandler == nil {
		unregistered = append(unregistered, "services.DeleteServiceHandler")
	}

	if o.ServicesGetServiceHandler == nil {
		unregistered = append(unregistered, "services.GetServiceHandler")
	}

	if o.ServicesListServicesHandler == nil {
		unregistered = append(unregistered, "services.ListServicesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the an application for easier distributed application generation API
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/services/{name}"] = services.NewCreateOrUpdateService(o.context, o.ServicesCreateOrUpdateServiceHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/services/{name}"] = services.NewDeleteService(o.context, o.ServicesDeleteServiceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{name}"] = services.NewGetService(o.context, o.ServicesGetServiceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services"] = services.NewListServices(o.context, o.ServicesListServicesHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *AnApplicationForEasierDistributedApplicationGenerationAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
