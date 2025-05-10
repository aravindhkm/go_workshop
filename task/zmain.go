package task

import "fmt"

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
