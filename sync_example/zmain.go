package syncexample

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	CondOne,
	Cond_Two,
	Cond_Two_One,
	CondThree,
	Cond_D_One,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./sync_example"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
