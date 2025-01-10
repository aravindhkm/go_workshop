package ctx

import (
	"context"
	"fmt"
	// "net/http"
	"time"
)

func Diff() {
	// context.Background() example
	fmt.Println("Example with context.Background():")
	backgroundCtx := context.Background()
	doSomething(backgroundCtx)

	// context.TODO() example
	fmt.Println("Example with context.TODO():")
	todoCtx := context.TODO()
	doSomething(todoCtx)
}

func doSomething(ctx context.Context) {
	// Simulating some work with context
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work done")
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	}
}