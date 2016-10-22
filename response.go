package grout

import "net/http"

// Response definition adding a convenience layer to the http.ResponseWriter
type Response struct {
	// Http status code issued to the cliend
	Status int

	// Header fields
	Header map[string]string

	rawWriter     http.ResponseWriter
	headerWritten bool
}

func (res *Response) Write(data []byte) (int, error) {
	if !res.headerWritten {
		res.writeHeaderData()
	}
	return res.rawWriter.Write(data)
}

func (res *Response) writeHeaderData() {
	for key, val := range res.Header {
		res.rawWriter.Header().Add(key, val)
	}
	res.rawWriter.WriteHeader(res.Status)
}

//
func (res *Response) SetCookie(c *http.Cookie) {
	http.SetCookie(res.rawWriter, c)
}

// NewResponse creates a new Response object and sets the status to 200
func NewResponse(rw http.ResponseWriter) Response {
	return Response{Status: 200, Header: make(map[string]string), rawWriter: rw}
}
