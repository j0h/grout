package gorouter

import (
	"log"
	"time"
)

func printLog(method, uri string, duration time.Duration, status int) {
	log.Printf(
		"[%s - %d]\t%s\t%s",
		method,
		status,
		uri,
		duration,
	)
}
