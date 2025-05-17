package interfacehandle

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	ItrFive,
	ItrFour,
	ItrOne,
	ItrOwnWork,
	ItrTwo,
	ItrThree,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./interface_handle"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
