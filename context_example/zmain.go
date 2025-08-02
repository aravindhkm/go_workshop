package contextexample

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	AfterFunc_A,
	ToDo_A,
	ToDo_B,
	WithCancel_A,
	WithCancel_B,
	WithDeadline_A,
	WithTimeout_A,
	WithValue_A,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./context_example"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}

// The context package in Go provides a way to carry deadlines,
// cancellation signals, and request-scoped values across API boundaries
// and between goroutines.

// It’s mainly used to:

// Control cancellation of goroutines (e.g., stop work if a client disconnects)

// Set deadlines or timeouts for operations

// Pass request-scoped data through function calls
