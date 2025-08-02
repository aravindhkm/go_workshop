package ratelimit

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	RateLimitOne,
	RateExampleTwo,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./rate_limit"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
