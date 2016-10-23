package grout

import "net/http"

// Response definition adding a convenience layer to the http.ResponseWriter. However, we do not
// implement the http.ResponseWriter interface on purpose as there are no convenient operations
// for setting the http Status for example. Not implementing the interface forces the use of the
// Response datastructure. Of course due do this, there are missing convenience operations. Please
// submit a pull request or an issue if there is something missing.
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
		res.headerWritten = true
	}
	return res.rawWriter.Write(data)
}

func (res *Response) writeHeaderData() {
	for key, val := range res.Header {
		res.rawWriter.Header().Add(key, val)
	}
	res.rawWriter.WriteHeader(res.Status)
}

// SetCookie sets a cookie on the underlying ResponseWriter.
func (res *Response) SetCookie(c *http.Cookie) {
	http.SetCookie(res.rawWriter, c)
}

// NewResponse creates a new Response object and sets the status to 200
func NewResponse(rw http.ResponseWriter) Response {
	return Response{Status: 200, Header: make(map[string]string), rawWriter: rw}
}
