package ctx

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

// context.WithDeadline(parent, deadline) creates a new context that
// will be automatically canceled at the specified time.

// defer cancel() ensures that resources are cleaned up even if the
// deadline isn't hit.

// ctx.Done() is a channel that is closed when the deadline expires
// or the context is canceled manually.

// ctx.Err() gives the reason for the cancellation (DeadlineExceeded
// or Canceled).
