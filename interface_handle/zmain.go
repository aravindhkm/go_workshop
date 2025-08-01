package interfacehandle

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	InterfaceBasicExample,
	InterfaceEmbed,
	ItrFive,
	ItrNilOne,
	ItrNilTwo,
	ItrOne,
	ItrOwnWork,
	ItrSix,
	ItrThree,
	ItrTypeAssign,
	ItrTypeFind,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./interface_handle"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	// fmt.Println("runKey", runKey)
	fn := funcNames[runKey]
	fn()
}
