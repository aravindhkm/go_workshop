package zexecute

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(string){
	AddCmd,
}

func Run(funcName string) {
	currDir := "./keyword"
	runKey, err := config.FindFuncKey(currDir, "1")

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}

	fn := funcNames[runKey]
	fn(funcName)
}
