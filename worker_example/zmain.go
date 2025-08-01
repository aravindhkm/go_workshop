package workerexample

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	AtomicExampleOne,
	CounterSemaphoreOne,
	CounterSemaphorePkg,
	WorkerPoolOne,
	WorkerPoolThree,
	WorkerPoolTwo,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./worker_example"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
