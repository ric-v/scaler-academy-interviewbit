package main

import (
	"net/http"
	_ "net/http/pprof" // This import registers pprof handlers with the default mux
	"time"
)

func worker() {
	// This goroutine does nothing but sleep
	time.Sleep(10 * time.Hour)
}

func main() {
	// Start a large number of goroutines
	for i := 0; i < 10000www; i++ {
		go worker()
	}

	// Start an HTTP server with default mux, which includes pprof handlers
	http.ListenAndServe(":8080", nil)
}
