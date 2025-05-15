package interview

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	FindPairSumsOne,
	SecondLargetNumberOne,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./interview"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
