package problem

import (
	arraystring "WorkShop/problem/array_string"
	"WorkShop/problem/linkedlist"
)

var execute string = "linkedlist"

func Run(funcName string) {
	if execute == "array_string" {
		arraystring.ArrayManipulationMain()
	}

	if execute == "linkedlist" {
		linkedlist.LinkedListMain()
	}
}
