package leetcode

import (
	"WorkShop/leetcode/bufio"                                     // bufio
	"WorkShop/leetcode/linkedlist"                                // linkedlist
	shufflearray "WorkShop/leetcode/shuffle_array"                // bufio
	sortingandsearching "WorkShop/leetcode/sorting-and-searching" // sortingandsearching
)

var exec string = "shufflearray"

func Run(funcName string) {
	if exec == "sortingandsearching" {
		sortingandsearching.SortingAndSearchingExecute()
	}

	if exec == "linkedlist" {
		linkedlist.LinkedListExecute()
	}

	if exec == "bufio" {
		bufio.BufioExecute()
	}

	if exec == "shufflearray" {
		shufflearray.ShuffleArrayMain()
	}
}
