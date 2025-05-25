package syncexample

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	Atomic_A,
	Atomic_B,
	Atomic_C,
	CondOne,
	Cond_Two,
	Cond_Two_One,
	CondThree,
	Cond_D_One,
	Map_A,
	Map_B,
	Pool_A,
	RW_Mutex_A,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./sync_example"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
