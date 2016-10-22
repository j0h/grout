package grout

// RouteDecorator wrapping the an initial handler. This way the handlers get decorated and we can provide more information
type RouteDecorator func(handler RouteHandler, r *Route) RouteHandler

// RouteHandler for a Route
type RouteHandler interface {
	Run(*Request, *Response)
}

// RouteHandlerFunc as specified in the RouteHandler interface
type RouteHandlerFunc func(*Request, *Response)

//Run implements the RouteHandler interface
func (h RouteHandlerFunc) Run(req *Request, res *Response) {
	h(req, res)
}

// Route definition
type Route struct {
	name    string
	methods []string
	pattern string
	handler RouteHandler
}

// SetName sets the route name
func (r *Route) SetName(name string) *Route {
	r.name = name
	return r
}

// GetName returns the route name
func (r *Route) GetName() string {
	return r.name
}

// SetMethods sets the list of http methods which can be used to call exactly this route
func (r *Route) SetMethods(methods ...string) *Route {
	r.methods = methods
	return r
}

// GetMethods returns the allowed http methods
func (r *Route) GetMethods() []string {
	return r.methods
}

// SetPattern sets the url pattern that matches the route
func (r *Route) SetPattern(pattern string) *Route {
	r.pattern = pattern
	return r
}

// GetPattern returns the raw string pattern as defined by the user and as it is used to match against an url
func (r *Route) GetPattern() string {
	return r.pattern
}

// SetHandler sets the Handler for the request
func (r *Route) SetHandler(handler RouteHandler) *Route {
	r.handler = handler
	return r
}

// GetHandler returns the Handler for the request
func (r *Route) GetHandler() RouteHandler {
	return r.handler
}
