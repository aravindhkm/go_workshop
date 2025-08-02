package zexecute

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(string){
	AddCmd,
}

func Run(funcName string) {
	currDir := "./zexecute"
	_, err := config.FindRunIndex(currDir, "1")

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}

	fn := funcNames[0]
	fn(funcName)
}
