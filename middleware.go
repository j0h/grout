package grout

import (
	"errors"
	"net/http"
)

// MiddlewareHandler is a function manipulating a request. If an error occures (err != nil) it is sent to the client. Useful for e.g. authentication
// It is guaranteed that r is never nil. This is the case that the route does not exist and thus an 404 is issued before.
type MiddlewareHandler func(req *http.Request, res *Response, route *Route) error

// Middleware holds meta data useful for debugging as well as the handler which is ought to be called. All middlewares run before executing the
type Middleware struct {
	id      int
	name    string
	handler MiddlewareHandler
}

// AddMiddleware adds an middleware to the set of active middlewares intercepting an incoming request. The middleware is returned afterwards
// It is not possible to access the internally used middlewares. If there is a usecase for this, support for this can be implemented fastly.
func (r *Router) AddMiddleware(name string, mw MiddlewareHandler) *Middleware {
	mid := &Middleware{id: r.middlewareIDCounter, name: name, handler: mw}
	r.activeMiddlewares[r.middlewareIDCounter] = mid
	return mid
}

// -- internal middlewares

func checkRouteValidity(req *http.Request, res *Response, route *Route) error {
	if route == nil {
		// issue 404
		res.Status = 404
		return errors.New("Could not find resource.")
	}
	return nil
}
