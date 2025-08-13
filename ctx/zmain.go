package ctx

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	AfterFunc_A,
	CtxWithValueExampleOne,
	Diff,
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
	currDir := "./ctx"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
