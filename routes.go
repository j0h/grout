package gorouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RouteDecorator wrapping the an initial handler. This way the handlers get decorated and we can provide more information
type RouteDecorator func(handler http.Handler, r Route) http.Handler

var routeDecorators = []RouteDecorator{}

// Route definition
type Route struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
	innerROute  *mux.Route
}

// CreateRoute creates a new Route
func CreateRoute(name, method, pattern string, handlerFunc http.HandlerFunc) *Route {
	return &Route{name: name, method: method, pattern: pattern, handlerFunc: handlerFunc}
}

// Add r to the list of available routes. r is online afterwards. This is not threadsafe!
func (r *Route) Add() {
	registeredRoutes = append(registeredRoutes, *r)
	routesChanged = true
}

//
func (r *Route) SetName(name string) {
	r.name = name
}

//
func (r *Route) GetName() string {
	return r.name
}

//
func (r *Route) SetMethod(method string) {
	r.method = method
}

//
func (r *Route) GetMethod() string {
	return r.method
}

//
func (r *Route) SetPattern(pattern string) {
	r.pattern = pattern
}

//
func (r *Route) GetPattern() string {
	return r.pattern
}

//
func (r *Route) SetHandlerFunc(handlerFunc http.HandlerFunc) {
	r.handlerFunc = handlerFunc
}

//
func (r *Route) GetHandlerFunc() *http.HandlerFunc {
	return &r.handlerFunc
}

var registeredRoutes = []Route{}
var routesChanged = true
