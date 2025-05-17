package httphandle

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	HttpGetOne,
	HttpPostExampleOne,
	HttpPostExampleTwo,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./http_handle"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
