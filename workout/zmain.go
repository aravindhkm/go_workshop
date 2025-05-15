package workout

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	DummyWorkOut,
	RaceCondition,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./workout"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
