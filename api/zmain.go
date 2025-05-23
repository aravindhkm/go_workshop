package api

import (
	"WorkShop/config"
	"fmt"
)

const MyExportedVar = "I am accessible"

var funcNames = []func(){
	GinExampleOne,
	GinExampleTwo,
	HttpApiOne,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./api"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
