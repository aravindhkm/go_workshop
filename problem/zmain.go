package problem

import (
	arraystring "WorkShop/problem/array_string"
	"WorkShop/problem/linkedlist"
)

var execute string = "linkedlist"

func Run(funcName string) {
	if execute == "array_string" {
		arraystring.ArrayManipulationMain()
	}

	if execute == "linkedlist" {
		linkedlist.LinkedListMain()
	}
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
