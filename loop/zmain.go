package loop

import (
	"WorkShop/loop/selectcase"
	"WorkShop/loop/switchcase"
	"fmt"
)

const context = "switchcase"

func Run(funcName string) {

	if context == "selectcase" {
		selectcase.SelectCaseExecute()
	}

	if context == "switchcase" {
		switchcase.SwitchExecute()
	}

	fmt.Println("")
}

import "fmt"

var funcNames = []func(){
	SimpleErrorHandle,
}

// Run executes a function by key from the store
func Run(funcName string) {
	var funcStore = map[string]func(){}

	// Populate the map with function names as keys
	for index, fn := range funcNames {
		funcStore[fmt.Sprintf("%d", index+1)] = fn
	}
	if fn, exists := funcStore[funcName]; exists {
		fn()
		return
	} else {
		funcStore["1"]()
		return
	}
}
