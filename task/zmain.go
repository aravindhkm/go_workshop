package task

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	PrintAddEvenGoRoutineMain,
	PrintAddEvenGoRoutine,
	AddEvenLiner,
	UnionIntersectExecute,
	FindMissingNumberInUnsortedArray,
	GarbageExecute,
	InventoryExecute,
	RemoveZeroFromArray,
	SelectFibonacciExample,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./task"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
