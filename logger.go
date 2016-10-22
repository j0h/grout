package gorouter

import (
	"log"
	"net/http"
	"time"
)

func printLog(req http.Request, res Response, duration time.Duration) {
	log.Printf(
		"[%s - %d]\t%s\t%s",
		req.Method,
		res.Status,
		req.RequestURI,
		duration,
	)
}
