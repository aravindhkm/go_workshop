package concurrency

// Store function mappings
var funcStore = map[string]func(){
	"1": BufferChannelOne,
	"2": BufferChannelTwo,
	"3": ChannelDirection,
	"4": GR_CH_Simple,
	"5": Gr_Simple_Sleep,
	"6": Gr_Simple_Sleep_Two,
	"7": ParallelismOne,
	"8": ParallelismExampleTwo,
	"9": ParallelismExampleThree,
	"10": WeatherApiGR,
	"11": WeatherApiNormal,
	"12": WgLockUnlockExampleOne,
}

// Run executes a function by key from the store
func Run(funcName string) {
	if fn, exists := funcStore[funcName]; exists {
		fn()
		return
	} else {
		funcStore["1"]()
		return
	}
}