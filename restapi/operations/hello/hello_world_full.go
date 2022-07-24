// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// HelloWorldFullHandlerFunc turns a function with the right signature into a hello world full handler
type HelloWorldFullHandlerFunc func(HelloWorldFullParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HelloWorldFullHandlerFunc) Handle(params HelloWorldFullParams) middleware.Responder {
	return fn(params)
}

// HelloWorldFullHandler interface for that can handle valid hello world full params
type HelloWorldFullHandler interface {
	Handle(HelloWorldFullParams) middleware.Responder
}

// NewHelloWorldFull creates a new http.Handler for the hello world full operation
func NewHelloWorldFull(ctx *middleware.Context, handler HelloWorldFullHandler) *HelloWorldFull {
	return &HelloWorldFull{Context: ctx, Handler: handler}
}

/* HelloWorldFull swagger:route GET /hello-world Hello World helloWorldFull

Example Route

Some description

*/
type HelloWorldFull struct {
	Context *middleware.Context
	Handler HelloWorldFullHandler
}

func (o *HelloWorldFull) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewHelloWorldFullParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
