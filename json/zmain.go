package json

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	JsonApi,
	ReadJsonFile,
	ReadJsonFileTwo,
	WriteJsonExampleOne,
	WriteJsonExampleTwo,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./json"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
