package contextexample

import (
	"context"
	"fmt"
	"time"
)

func doWorkDeadline(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second): // Simulates a long-running task
		return nil // Task completed
	case <-ctx.Done(): // Context deadline exceeded or cancelled
		return ctx.Err() // Returns context.DeadlineExceeded or context.Canceled
	}
}

func WithDeadline_A() {
	// Set a deadline 2 seconds from now
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // Ensure the context is cancelled to release resources

	// Start a task that takes 3 seconds
	err := doWorkDeadline(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}
