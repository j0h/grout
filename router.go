package grout

import (
	"net/http"
	"time"
)

var supportedMethods = []string{"GET", "POST", "PUT", "DELETE", "UPDATE"}

// Router takes care of matching the routes as well as attaching and executing middlewares/decorators.
type Router struct {
	routeDecorators []RouteDecorator

	registeredRoutes  map[string][]*Route
	activeMiddlewares map[int]*Middleware

	middlewareIDCounter int

	RouteMatcher Matcher
}

// NewRouter creates a new router handling routes.
func NewRouter() *Router {
	router := &Router{
		activeMiddlewares: make(map[int]*Middleware),
		registeredRoutes:  make(map[string][]*Route),
		RouteMatcher:      getDefaultMatcher(),
	}
	router.AddMiddleware("_grout_routeValidityCheck__", checkRouteValidity)

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
	route, match := r.GetRouteByPath(req.URL.Path, req.Method)
	request := convertToRequest(req, *match)

	var err error
	for _, middleware := range r.activeMiddlewares {
		err = middleware.handler(&request, &res, route)
		if err != nil {
			// send error to client
			res.Write([]byte(err.Error()))
			break
		}
	}

	if err != nil && res.Status == 200 {
		res.Status = 500
	}

	if match != nil && err == nil {
		route.handler.Run(&request, &res)
	}

	printLog(*req, res, time.Since(start))
}

// GetRouteByPath and match it against included patterns
func (r *Router) GetRouteByPath(path, method string) (*Route, *MatchResult) {
	for _, route := range r.registeredRoutes[method] {
		if match := r.RouteMatcher.Match(path, route); match != nil {
			return route, match
		}
	}
	return nil, nil
}

// NewRoute returns a blank route. The route is not added to the set of active routes, yet.
func (r *Router) NewRoute() *Route {
	return &Route{}
}

// AddRoute r to the list of available routes. r is online afterwards.
func (r *Router) AddRoute(route *Route, methods ...string) {
	for _, m := range methods {
		r.registeredRoutes[m] = append(r.registeredRoutes[m], route)

		// apply current decorators
		for _, decorator := range r.routeDecorators {
			route.SetHandler(decorator(route.GetHandler(), route))
		}
	}
}

// CreateRoute creates a new Route and adds it to the router
func (r *Router) CreateRoute(name, pattern string, handlerFunc RouteHandlerFunc, methods ...string) *Route {
	newRoute := &Route{name: name, methods: methods, pattern: pattern, handler: handlerFunc}
	r.AddRoute(newRoute, methods...)
	return newRoute
}

// AddRouteDecorator and apply it to existing routes
func (r *Router) AddRouteDecorator(decorator RouteDecorator) {
	r.routeDecorators = append(r.routeDecorators, decorator)
	// apply it to all current routes
	for _, m := range supportedMethods {
		for _, route := range r.registeredRoutes[m] {
			route.SetHandler(decorator(route.GetHandler(), route))
		}
	}
}

func convertToRequest(req *http.Request, match MatchResult) Request {
	return Request{
		Body:        req.Body,
		MatchResult: match,
		HTTPRequest: req,
	}
}
