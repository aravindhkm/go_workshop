package contextexample

import (
	"context"
	"fmt"
)

type contextKey string

func getCtxValue(ctx context.Context, userIDKey contextKey) {
	if data := ctx.Value(userIDKey); data != nil {
		fmt.Println("Result", "hello", data)
	}
}
func WithValue_A() {
	userIDKey := contextKey("userID")

	ctx := context.WithValue(context.Background(), userIDKey, 10)

	getCtxValue(ctx, userIDKey)
}
