package gorouter

import (
	"log"
	"net/http"
	"time"
)

// LoggerRouteDecorator is a basic decoration of a route. Is is an example as well as a useful tool.
// Add this to the set of decorators if you want to have a log of the form
// [<Method> - <StatusCode>] <Name>:<RequestURI> <elapsed time>
func LoggerRouteDecorator(innerHandler RouteHandler, route *Route) RouteHandler {
	return RouteHandlerFunc(func(req *http.Request, res *Response) {
		start := time.Now()

		innerHandler.Run(req, res)

		log.Printf(
			"[%s - %d]\t%s\t%s",
			req.Method,
			res.Status,
			req.RequestURI,
			time.Since(start),
		)
	})
}
