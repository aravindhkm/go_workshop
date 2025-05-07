package ctx

// Store function mappings
var funcStore = map[string]func(){
	"1": CtxWithValueExampleOne,
	"2": Diff,
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