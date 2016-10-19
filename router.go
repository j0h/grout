package gorouter

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

//
type Router struct {
	router *mux.Router
}

// NewRouter creates a new router handling routes.
func NewRouter() *Router {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := &Router{router: muxRouter}
	router.ReloadRoutes()

	return router
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*") // available to everyone
	route := r.GetRouteByPath(req.URL.Path)
	if route == nil {
		rw.WriteHeader(404)
		rw.Write([]byte("404"))
		return
	}

	for _, middleware := range activeMiddlewares {
		err := (*middleware).handler(req, route)
		if err != nil {
			// send error to client
			rw.Write([]byte(err.Error()))
			return
		}
	}
	r.router.ServeHTTP(rw, req)
}

//
func (r *Router) GetRouteByPath(path string) *Route {
	for _, route := range registeredRoutes {
		if matched, _ := regexp.Match(route.GetPattern(), []byte(path)); matched {
			return &route
		}
	}
	return nil
}

//
func (r *Router) ReloadRoutes() {
	if !routesChanged {
		return
	}

	log.Println("Loading routes...")

	for _, route := range registeredRoutes {
		internalRoute := r.router.GetRoute(route.GetName())
		var handler http.Handler = route.GetHandlerFunc() // implicit interfaces are confusing

		for _, h := range routeDecorators {
			handler = h(handler, route)
		}

		if internalRoute == nil {
			// this route is new, register new one
			r.router.
				Methods(route.GetMethod()).
				Path(route.GetPattern()).
				Name(route.GetName()).
				Handler(route.GetHandlerFunc())
		} else {
			// this route is existing, adjust it - however, it is not clear whether these changes have any influence on the underlying routing
			internalRoute.
				Name(route.GetName()).
				Methods(route.GetMethod()).
				Path(route.GetPattern()).
				Handler(route.GetHandlerFunc())
		}
	}

	routesChanged = false
}
