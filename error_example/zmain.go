package errorexample

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	OwnErrorMsgOne,
	SimpleErrorHandle,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./error_example"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
