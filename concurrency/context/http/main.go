package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// A slow computation function
func slowComputation(ctx context.Context) (int, error) {
	// Simulate a slow computation
	select {
	case <-time.After(5 * time.Second): // Assume computation takes 5 seconds
		return 42, nil // Return the computed value
	case <-ctx.Done():
		return 0, ctx.Err() // If the context is cancelled, return an error
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := slowComputation(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintf(w, "Computation result: %v\n", result)
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
