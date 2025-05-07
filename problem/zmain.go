package problem

import (
	arraystring "WorkShop/problem/array_string"
	"WorkShop/problem/linkedlist"
)

var floderName = map[string]func(string){
	"array_string": arraystring.Run,
	"linkedlist": linkedlist.Run,
}

var execFloder string = "shufflearray"

// Run executes a function by key from the store
func Run(funcName string) {	
	if fn, exists := floderName[execFloder]; exists {
		fn(funcName)
		return
	} else {
		floderName[execFloder](funcName)
		return
	}
}
