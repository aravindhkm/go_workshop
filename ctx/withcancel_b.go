package ctx

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// handler simulates some work and respects context cancellation.
func handler(w http.ResponseWriter, r *http.Request) {
	// Create a cancellable context derived from the request's context
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Listen for client disconnect in a separate goroutine
	go func() {
		<-ctx.Done() // This will unblock if the context is canceled
		log.Println("Request was cancelled by the client")
	}()

	// Simulate some work that checks the context cancellation
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			// Context was canceled; stop work and return
			http.Error(w, "Request canceled", http.StatusRequestTimeout)
			return
		case <-time.After(1 * time.Second):
			// Simulating work by sleeping 1 second
			log.Printf("Processing step %d\n", i+1)
		}
	}

	// Write response if work finished successfully
	fmt.Fprintln(w, "Work completed successfully")
}

func WithCancel_B() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
	log.Println("Server gracefully stopped")
}
