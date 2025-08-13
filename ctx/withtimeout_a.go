package ctx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// fetchData performs an HTTP GET with a context timeout.
func fetchData(ctx context.Context, url string) error {
	// Create a new HTTP request with the context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// HTTP client with default settings
	client := &http.Client{}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// Check if the error is due to the context deadline
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("request timed out")
		}
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Optional: Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println("Request successful")
	return nil
}
func WithTimeout_A() {
	// Set a timeout duration
	timeout := 2 * time.Second

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // Always call cancel to release resources

	// Call the function with the context
	url := "https://httpbin.org/delay/3" // Delays response for 3 seconds
	if err := fetchData(ctx, url); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// | Feature          | `WithTimeout`                  | `WithDeadline`                 |
// | ---------------- | ------------------------------ | ------------------------------ |
// | Timeout type     | Relative (e.g., 5s from now)   | Absolute (e.g., at 3:15:00 PM) |
// | Typical use case | API calls, retries, DB queries | Sync with external deadline    |
// | Argument         | Duration (`time.Duration`)     | Absolute time (`time.Time`)    |
// | Internally uses  | `WithDeadline`                 | Native deadline handling       |

// Use **WithTimeout** when you just want a task to timeout after a set period.

// Use **WithDeadline** when aligning with a fixed clock time, like scheduling or external API requirements.
