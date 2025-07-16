package interview

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	AnagramOne,
	BiMapOne,
	BiMapTwo,
	FacOne,
	FibTimePing,
	Fibonacci,
	FindChar,
	FindCharDuplicate,
	FindPairSumsOne,
	LinkedListExec,
	Palindrome_One,
	PianoOne,
	ReverseIntAndSum,
	SecondLargetNumberOne,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./interview"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}

/*
./run.sh interview
./cmd.sh interview
*/
