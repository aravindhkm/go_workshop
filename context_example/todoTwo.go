package contextexample

import (
	"context"
	"fmt"
	"time"
)

// ❌ This handler uses context.TODO(), which does nothing on cancel/timeout
func handlerUsingTODOTwo() {
	ctx := context.Background()
	doWorkTwo(ctx) // Work doesn't cancel even if the client disconnects
	fmt.Println("Finished with TODO context (won't stop early)")
	fmt.Println()
}

// ✅ This handler uses a timeout to cancel slow work
func handlerWithTimeoutTwo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	doWorkTwo(ctx) // Work will be cancelled if it exceeds 2s
	fmt.Println("Finished with timeout context (2s max)")
	fmt.Println()

}

// Simulated long-running task that respects context cancellation
func doWorkTwo(ctx context.Context) {
	fmt.Println("Starting work...")
	select {
	case <-time.After(5 * time.Second): // Simulate a long task
		fmt.Println("Work complete")
	case <-ctx.Done(): // Stop if context is cancelled or times out
		fmt.Println("Work cancelled:", ctx.Err())
	}
}

func ToDoTwo() {
	handlerUsingTODOTwo()
	handlerWithTimeoutTwo()
}
