// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"

	"practiceproject/restapi/operations"
	"practiceproject/restapi/operations/hello"
)

//go:generate swagger generate server --target ..\..\practiceproject --name ExampleService --spec ..\swagger-api\swagger.yml --principal interface{}

func ConfigureFlags(api *operations.ExampleServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func ConfigureAPI(api *operations.ExampleServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = JSONConsumer()

	api.JSONProducer = JSONProducer()

	if api.HelloHelloWorldHandler == nil {
		api.HelloHelloWorldHandler = hello.HelloWorldHandlerFunc(func(params hello.HelloWorldParams) middleware.Responder {
			return middleware.NotImplemented("operation hello.HelloWorld has not yet been implemented")
		})
	}
	if api.HelloHelloWorldFullHandler == nil {
		api.HelloHelloWorldFullHandler = hello.HelloWorldFullHandlerFunc(func(params hello.HelloWorldFullParams) middleware.Responder {
			return middleware.NotImplemented("operation hello.HelloWorldFull has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return SetupGlobalMiddleware(api.Serve(SetupMiddlewares))
}

func ServeError(writer http.ResponseWriter, request *http.Request, err error) {

}

func JSONConsumer() runtime.Consumer {
	return nil
}

func JSONProducer() runtime.Producer {

	return nil
}

// The TLS configuration before HTTPS server starts.
func ConfigureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func ConfigureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
