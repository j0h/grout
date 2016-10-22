package gorouter

import (
	"io"
	"net/http"
)

// Request is holding all information like parameter matching and body content
type Request struct {
	Body        io.ReadCloser
	MatchResult MatchResult
	httpRequest *http.Request
}
