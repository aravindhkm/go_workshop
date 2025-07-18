package api

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// HttpApiOne starts the HTTP server and listens for shutdown signals.
func HttpApiOne() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()

	// Register endpoint
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Pong"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		fmt.Println("Starting server on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	fmt.Println("HTTP API is running successfully")

	// Wait for context cancellation (signal interrupt)
	<-ctx.Done()
	fmt.Println("\nShutting down server...")

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	} else {
		fmt.Println("Server stopped gracefully.")
	}
}
