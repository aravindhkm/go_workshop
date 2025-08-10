package keyword

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	AnyOne,
	ReferenceOrValue,
	ArrayTest,
	ByteOne,
	CapOne,
	Copy_A,
	Copy_B,
	Copy_C,
	DeferExecute,
	ForRange,
	IotaExampleOne,
	MakeVsNewExampleOne,
	MapKey,
	NilOne,
	PtrOne,
	PtrTwo,
	RecoverExampleOne,
	RecoverChTwo,
	ReferenceOne,
	RuneExampleOne,
	SelectOne,
	SliceTest,
	SliceAndCap,
	StringImmutable,
	StructExampleOne,
	SwapTwoNumberExecute,
	TypeAll,
	TypeInt,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./keyword"
	runKey, err := config.FindRunIndex(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}

/*

./cmd.sh keyword

*/
