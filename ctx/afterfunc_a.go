package ctx

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWorkAftFn(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Printf("Work interrupted by context: %v", ctx.Err())
	case <-time.After(5 * time.Second):
		log.Println("Work completed successfully.")
	}
}

func cleanup() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Resources cleaned up.")
}

func AfterFunc_A() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	done := make(chan struct{}) // used to wait for AfterFunc cleanup

	context.AfterFunc(ctx, func() {
		log.Println("Context was canceled. Running cleanup task.")
		cleanup()
		close(done) // signal that cleanup is done
	})

	log.Println("Starting work...")
	doWorkAftFn(ctx)

	<-done // wait for AfterFunc to complete
	log.Println("Main function completed.")
}
