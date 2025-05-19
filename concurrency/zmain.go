package concurrency

import (
	"WorkShop/config"
	"fmt"
)

var funcNames = []func(){
	BufferChannelOne,
	BufferChannelTwo,
	ChannelDirection,
	CloseChOne,
	GR_CH_Simple,
	Gr_Simple_Sleep,
	Gr_Simple_Sleep_Two,
	ParallelismOne,
	ParallelismExampleTwo,
	ParallelismExampleThree,
	WeatherApiGR,
	WeatherApiNormal,
	WgLockUnlockExampleOne,
}

// Run executes a function by key from the store
func Run(funcName string) {
	currDir := "./concurrency"
	runKey, err := config.FindFuncKey(currDir, funcName)

	if err != nil {
		fmt.Printf("Error in %s: %s\n", currDir, err.Error())
	}
	fn := funcNames[runKey]
	fn()
}
