package leetcode

import (
	"WorkShop/leetcode/linkedlist"  // linkedlist
	"WorkShop/leetcode/sorting-and-searching"  // sortingandsearching
	"WorkShop/leetcode/bufio"  // bufio
	"WorkShop/leetcode/shuffle_array"  // bufio
)

var exec string = "shufflearray"

func LeetcodeExecute() {
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