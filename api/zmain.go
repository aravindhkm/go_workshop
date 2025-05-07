package api

// Store function mappings
var funcStore = map[string]func(){
	"1": GinExampleOne,
	"2": GinExampleTwo,
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
