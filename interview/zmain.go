package interview

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	AnagramOne,
	BiMapOne,
	BiMapTwo,
	ClockAngle,
	CountWord,
	DeadLockOne,
	FacOne,
	FibTimePing,
	Fibonacci,
	FindChar,
	FindCharDuplicate,
	FindPairSumsOne,
	GenericChOne,
	GenericChTwo,
	InheritenceOne,
	LinkedListExec,
	OddEvenFind,
	Palindrome_One,
	PianoOne,
	ReverseIntAndSum,
	SecondLargetNumberOne,
	SignalOne,
	SignalTwo,
	SqrtSort,
	TwoSumArrOne,
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
