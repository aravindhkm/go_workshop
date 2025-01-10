package problem

import (
	"WorkShop/problem/array_string"
	"WorkShop/problem/linkedlist"
)

var execute string = "linkedlist"

func ProblemExecute() {
	if execute == "array_string" {
		arraystring.ArrayManipulationMain()
	}

	if execute == "linkedlist" {
		linkedlist.LinkedListMain()
	}
}