package gorouter

import "net/http"

// Response definition adding a convenience layer to the http.ResponseWriter
type Response struct {
	// Http status code issued to the cliend
	Status int

	// Header fields
	Header map[string]string

	rawWriter http.ResponseWriter
}

func (res *Response) Write(data []byte) (int, error) {
	return res.rawWriter.Write(data)
}

// write response to the client
func (res *Response) write() {
	res.writeHeader()
}

func (res *Response) writeHeader() {
	res.rawWriter.WriteHeader(res.Status)
	for key, val := range res.Header {
		res.rawWriter.Header().Add(key, val)
	}
}

// NewResponse creates a new Response object and sets the status to 200
func NewResponse(rw http.ResponseWriter) Response {
	return Response{Status: 200, Header: make(map[string]string), rawWriter: rw}
}
