package gorouter

import "net/http"

// MiddlewareHandler is a function manipulating a request. If an error occures (err != nil) it is send to the client. Useful for e.g. authentication
// It is guaranteed that r is never nil. This is the case that the route does not exist and thus an 404 is issued before.
type MiddlewareHandler func(req *http.Request, r *Route) error

// Middleware holds meta data useful for debugging as well as the handler which is ought to be called. All middlewares run before executing the
type Middleware struct {
	id      int
	name    string
	handler MiddlewareHandler
}

var activeMiddlewares = []*Middleware{}

// AddMiddleware adds an middleware to the set of active middlewares intercepting an incoming request. An id is returned.
func AddMiddleware(name string, mw MiddlewareHandler) int {
	activeMiddlewares = append(activeMiddlewares, &Middleware{id: nextID, name: name, handler: mw})
	nextID++
	return nextID - 1
}

var nextID = 0
