package task

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	PrintAddEvenGoRoutineMain,
	PrintAddEvenGoRoutine,
	AddEvenLinerOne,
	AddEvenLinerTwo,
	AddEvenLinerThree,
	AddEvenMoreOne,
	UnionIntersectExecute,
	FibonacciOne,
	FibonacciTwo,
	FindChar,
	FindMissingNumberInUnsortedArray,
	GarbageExecute,
	InventoryExecute,
	MergeSortOne,
	PrimeNumOne,
	ProfitStockOne,
	QueueOne,
	RemoveZeroFromArray,
	SelectFibonacciExample,
	StackOne,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./task"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
