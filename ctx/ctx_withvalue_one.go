package ctx

import (
	"context"
	"fmt"
)

type contextkey string

func processSensitiveData(ctx context.Context) {
	authToken, ok := ctx.Value(contextkey("AuthToken")).(string)
	if !ok {
		fmt.Println("Access Denied - missing AuthToken!")
		return
	}

	fmt.Printf("Processing data using: %s\n", authToken)
}

func CtxWithValueExampleOne() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, contextkey("AuthToken"), "X@cfs645JHSDcdae")

	processSensitiveData(ctx)
}
