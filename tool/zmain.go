package tool

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	GoCPU_Debug,
	VetExample,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./tool"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
