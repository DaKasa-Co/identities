// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/DaKasa-Co/identities/docs/restapi/operations"
)

//go:generate swagger generate server --target ../../docs --name IdentitiesMicrosservice --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.IdentitiesMicrosserviceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.IdentitiesMicrosserviceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-Key" header is set
	if api.APIKeyHeaderAuth == nil {
		api.APIKeyHeaderAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (APIKeyHeader) X-API-Key from header param [X-API-Key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostChallRecoveryMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostLoginMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostRecoveryMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostRegisterMaxParseMemory = 32 << 20

	if api.PostChallRecoveryHandler == nil {
		api.PostChallRecoveryHandler = operations.PostChallRecoveryHandlerFunc(func(params operations.PostChallRecoveryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostChallRecovery has not yet been implemented")
		})
	}
	if api.PostLoginHandler == nil {
		api.PostLoginHandler = operations.PostLoginHandlerFunc(func(params operations.PostLoginParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostLogin has not yet been implemented")
		})
	}
	if api.PostRecoveryHandler == nil {
		api.PostRecoveryHandler = operations.PostRecoveryHandlerFunc(func(params operations.PostRecoveryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostRecovery has not yet been implemented")
		})
	}
	if api.PostRegisterHandler == nil {
		api.PostRegisterHandler = operations.PostRegisterHandlerFunc(func(params operations.PostRegisterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostRegister has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
