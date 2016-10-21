package gorouter

import (
	"net/http"
	"time"
)

// Router takes care of matching the routes as well as attaching and executing middlewares/decorators.
type Router struct {
	routeDecorators []RouteDecorator

	registeredRoutes  []*Route
	activeMiddlewares map[int]*Middleware

	middlewareIDCounter int
}

// NewRouter creates a new router handling routes.
func NewRouter() *Router {
	router := &Router{activeMiddlewares: make(map[int]*Middleware)}
	router.AddMiddleware("__routeValidityCheck__", checkRouteValidity)

	return router
}

// Serve launches an http server listening on addr. If addr is empty it will listen on :http
// Attention: This is up to change for ssl support and probably some more...
func (r *Router) Serve(addr string) error {
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	return server.ListenAndServe()
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()

	res := NewResponse(rw)
	route := r.GetRouteByPath(req.URL.Path)
	var err error
	for _, middleware := range r.activeMiddlewares {
		err = middleware.handler(req, &res, route)
		if err != nil {
			// send error to client
			res.Write([]byte(err.Error()))
			break
		}
	}

	if err != nil && res.Status == 200 {
		res.Status = 500
	}

	res.write()
	if route != nil {
		route.handler.Run(req, &res)
	}

	printLog(req.Method, req.RequestURI, time.Since(start), res.Status)
}

// GetRouteByPath and match it against included patterns
func (r *Router) GetRouteByPath(path string) *Route {
	for _, route := range r.registeredRoutes {
		//if matched, _ := regexp.Match(route.GetPattern(), []byte(path)); matched {
		if route.GetPattern() == path {
			return route
		}
	}
	return nil
}

// NewRoute returns a blank route. The route is not added to the set of active routes, yet.
func (r *Router) NewRoute() *Route {
	return &Route{}
}

// AddRoute r to the list of available routes. r is online afterwards.
func (r *Router) AddRoute(route *Route) {
	r.registeredRoutes = append(r.registeredRoutes, route)
	// apply current decorators
	for _, decorator := range r.routeDecorators {
		route.SetHandler(decorator(route.GetHandler(), route))
	}
}

// CreateRoute creates a new Route and adds it to the router
func (r *Router) CreateRoute(name, pattern string, handlerFunc RouteHandlerFunc, methods ...string) *Route {
	newRoute := &Route{name: name, methods: methods, pattern: pattern, handler: handlerFunc}
	r.AddRoute(newRoute)
	return newRoute
}

// AddRouteDecorator and apply it to existing routes
func (r *Router) AddRouteDecorator(decorator RouteDecorator) {
	r.routeDecorators = append(r.routeDecorators, decorator)
	// apply it to all current routes
	for _, route := range r.registeredRoutes {
		route.SetHandler(decorator(route.GetHandler(), route))
	}
}
