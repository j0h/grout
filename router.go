package gorouter

import "net/http"

// Router takes care of matching the routes as well as attaching and executing middlewares/decorators.
type Router struct {
	routeDecorators []RouteDecorator

	registeredRoutes  []*Route
	activeMiddlewares []*Middleware
}

// NewRouter creates a new router handling routes.
func NewRouter() *Router {
	router := &Router{}

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
	rw.Header().Set("Access-Control-Allow-Origin", "*") // available to everyone
	route := r.GetRouteByPath(req.URL.Path)
	if route == nil {
		rw.WriteHeader(404)
		rw.Write([]byte("404"))
		return
	}

	for _, middleware := range r.activeMiddlewares {
		err := middleware.handler(req, route)
		if err != nil {
			// send error to client
			rw.Write([]byte(err.Error()))
			return
		}
	}
	route.handler.ServeHTTP(rw, req)
}

//
func (r *Router) GetRouteByPath(path string) *Route {
	for _, route := range r.registeredRoutes {
		//if matched, _ := regexp.Match(route.GetPattern(), []byte(path)); matched {
		if route.GetPattern() == path {
			return route
		}
	}
	return nil
}

//
func (r *Router) GetRoute(name string) *Route {
	return nil
}

// NewRoute returns a blank route. The route is not added to the set of active routes, yet.
func (r *Router) NewRoute() *Route {
	return &Route{}
}

// AddRoute r to the list of available routes. r is online afterwards. This is not threadsafe!
func (r *Router) AddRoute(route *Route) {
	r.registeredRoutes = append(r.registeredRoutes, route)
	// apply current decorators
	for _, decorator := range r.routeDecorators {
		route.SetHandler(decorator(route.GetHandler(), route))
	}
}

// CreateRoute creates a new Route and adds it to the router
func (r *Router) CreateRoute(name, pattern string, handlerFunc http.HandlerFunc, methods ...string) *Route {
	newRoute := &Route{name: name, methods: methods, pattern: pattern, handler: handlerFunc}
	r.AddRoute(newRoute)
	return newRoute
}

//
func (r *Router) AddRouteDecorator(decorator RouteDecorator) {
	r.routeDecorators = append(r.routeDecorators, decorator)
	// apply it to all current routes
	for _, route := range r.registeredRoutes {
		route.SetHandler(decorator(route.GetHandler(), route))
	}
}
