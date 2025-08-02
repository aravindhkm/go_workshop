package routine

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	SequenceMain,
	PingPong,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./routine"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
