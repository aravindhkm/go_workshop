package contextexample

import (
	"context"
	"fmt"
	"time"
)

func cancelMonitor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("Cancel Triggered")
				return

			}
		default:
			fmt.Println("doWork: working...")
		}
	}
}

func WithCancel_A() {
	ctx, cancel := context.WithCancel(context.Background())

	go cancelMonitor(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main: calling cancel()")
	cancel()
	fmt.Println("Main: end cancel()")

	time.Sleep(1 * time.Second)
	fmt.Println("Main: done")
}
