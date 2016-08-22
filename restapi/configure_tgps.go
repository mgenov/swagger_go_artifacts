package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/mgenov/temp/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name tgps --spec https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json

func configureFlags(api *operations.TgpsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TgpsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AddPetHandler = operations.AddPetHandlerFunc(func(params operations.AddPetParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddPet has not yet been implemented")
	})
	api.DeletePetHandler = operations.DeletePetHandlerFunc(func(params operations.DeletePetParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeletePet has not yet been implemented")
	})
	api.FindPetByIDHandler = operations.FindPetByIDHandlerFunc(func(params operations.FindPetByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindPetByID has not yet been implemented")
	})
	api.FindPetsHandler = operations.FindPetsHandlerFunc(func(params operations.FindPetsParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindPets has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
