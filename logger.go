package grout

import (
	"log"
	"net/http"
	"time"

	log15 "github.com/inconshreveable/log15"
)

// Log is a log15 instance which is the base for all loggers in grout
var Log = log15.New("module", "grout")

func printLog(req http.Request, res Response, duration time.Duration) {
	log.Printf(
		"[%s - %d]\t%s\t%s",
		req.Method,
		res.Status,
		req.RequestURI,
		duration,
	)
}
