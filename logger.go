package gorouter

import (
	"log"
	"net/http"
	"time"
)

// LoggerRouteDecorator is a basic decoration of a route. Is is an example as well as a useful tool.
// Add this to the set of decorators if you want to have a log of the form
// [<Method>] <Name>:<RequestURI> <elapsed time>
func LoggerRouteDecorator(innerHandler http.Handler, route Route) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		innerHandler.ServeHTTP(rw, r)

		log.Printf(
			"[%s]\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
