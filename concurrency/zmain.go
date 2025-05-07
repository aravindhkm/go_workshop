package concurrency

import "fmt"

var funcNames = []func(){
	BufferChannelOne,
	BufferChannelTwo,
	ChannelDirection,
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
