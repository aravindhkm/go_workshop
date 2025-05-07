package leetcode

import (
	"WorkShop/leetcode/bufio"                                     // bufio
	"WorkShop/leetcode/linkedlist"                                // linkedlist
	shufflearray "WorkShop/leetcode/shuffle_array"                // bufio
	sortingandsearching "WorkShop/leetcode/sorting-and-searching" // sortingandsearching
)

var floderName = map[string]func(string){
	"sortingandsearching": sortingandsearching.Run,
	"linkedlist": linkedlist.Run,
	"bufio": bufio.Run,
	"shufflearray": shufflearray.Run,
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


