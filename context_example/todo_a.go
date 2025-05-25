package contextexample

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ❌ This handler uses context.TODO(), which does nothing on cancel/timeout
func handlerUsingTODO(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	doWork(ctx) // Work doesn't cancel even if the client disconnects
	fmt.Fprintln(w, "Finished with TODO context (won't stop early)")
}

// ✅ This handler uses a timeout to cancel slow work
func handlerWithTimeout(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	doWork(ctx) // Work will be cancelled if it exceeds 2s
	fmt.Fprintln(w, "Finished with timeout context (2s max)")
}

// ✅ This handler uses the request context, which cancels if the client disconnects
func handlerWithRequestContext(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Automatically cancels if client closes the connection

	doWork(ctx)
	fmt.Fprintln(w, "Finished with request context (client-aware)")
}

// Simulated long-running task that respects context cancellation
func doWork(ctx context.Context) {
	fmt.Println("Starting work...")
	select {
	case <-time.After(5 * time.Second): // Simulate a long task
		fmt.Println("Work complete")
	case <-ctx.Done(): // Stop if context is cancelled or times out
		fmt.Println("Work cancelled:", ctx.Err())
	}
}

func ToDo_A() {
	http.HandleFunc("/todo", handlerUsingTODO)
	http.HandleFunc("/timeout", handlerWithTimeout)
	http.HandleFunc("/request", handlerWithRequestContext)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)

}
