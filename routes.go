package gorouter

import "net/http"

// RouteDecorator wrapping the an initial handler. This way the handlers get decorated and we can provide more information
type RouteDecorator func(handler http.Handler, r *Route) http.Handler

// Route definition
type Route struct {
	name    string
	methods []string
	pattern string
	handler http.Handler
}

//
func (r *Route) SetName(name string) *Route {
	r.name = name
	return r
}

//
func (r *Route) GetName() string {
	return r.name
}

//
func (r *Route) SetMethods(methods ...string) *Route {
	r.methods = methods
	return r
}

//
func (r *Route) GetMethods() []string {
	return r.methods
}

//
func (r *Route) SetPattern(pattern string) *Route {
	r.pattern = pattern
	return r
}

//
func (r *Route) GetPattern() string {
	return r.pattern
}

//
func (r *Route) SetHandler(handler http.Handler) *Route {
	r.handler = handler
	return r
}

//
func (r *Route) GetHandler() http.Handler {
	return r.handler
}
