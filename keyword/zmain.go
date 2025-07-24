package keyword

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	ReferenceOrValue,
	ArrayTest,
	ByteOne,
	Copy_A,
	Copy_B,
	Copy_C,
	DeferExecute,
	ForRange,
	IotaExampleOne,
	MakeVsNewExampleOne,
	MapKey,
	PtrOne,
	PtrTwo,
	RecoverCh,
	RecoverExampleOne,
	RuneExampleOne,
	SelectOne,
	SliceAndCap,
	SliceTest,
	StringImmutable,
	StructExampleOne,
	SwapTwoNumberExecute,
	TypeAll,
	TypeInt,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./keyword"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}

/*

./cmd.sh keyword

*/
