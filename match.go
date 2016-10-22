package gorouter

import "strings"

// Matcher describes a strategy for matching
type Matcher interface {
	// If the returned match is nil, there is no match to this route
	Match(uri string, r *Route) *MatchResult
}

// MatchResult is the result from matching an URL against a Matcher
type MatchResult struct {
	URLParams   map[string]string
	RouteParams map[string]string
}

//
type MatcherFunc func(uri string, r *Route) *MatchResult

//
func (matcher MatcherFunc) Match(uri string, r *Route) *MatchResult {
	return matcher(uri, r)
}

func getDefaultMatcher() Matcher {
	return MatcherFunc(DefaultMatcher)
}

// DefaultMatcher is up to change. It only detects url paramters and takes trailing slashes into account. A parameter is indicated by placing ":" in
// the route. A route which can be matched is for example /user/4 if there is a route /user/:id. Afterwards there will be a mapentry for id which maps to 4
func DefaultMatcher(uri string, r *Route) *MatchResult {
	// extract route components
	pattern := r.GetPattern()
	uriComponents := strings.Split(uri, "/")
	patternComponents := strings.Split(pattern, "/")

	result := &MatchResult{RouteParams: make(map[string]string), URLParams: make(map[string]string)}

	if len(uriComponents) != len(patternComponents) {
		return nil
	}

	for i, component := range uriComponents {
		patternComp := patternComponents[i]
		if strings.Index(patternComp, ":") == 0 {
			// routeParameter to match
			result.RouteParams[patternComp[1:]] = component
		} else {
			// no parameter to match, if this route matches, the components must be equal
			if patternComp != component {
				return nil
			}
		}
	}

	return result
}
