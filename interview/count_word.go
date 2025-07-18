package interview

import (
	"fmt"
	"strings"
)

func countWordOne(str string) {
	wordMap := map[string]int{}

	localString := ""
	for _, value := range str {
		if string(value) != " " {
			localString += string(value)
		} else {
			wordMap[localString]++
			localString = ""
		}
	}
	wordMap[localString]++

	for key, value := range wordMap {
		fmt.Println(key, value)
	}
}

func countWordTwo(str string) {
	res := strings.Split(str, " ")

	wordMap := map[string]int{}

	for _, value := range res {
		wordMap[value]++
		fmt.Println("split", value)
	}

	fmt.Println("__________________________")

	for key, value := range wordMap {
		fmt.Println(key, value)
	}
}

func countWordsValid(input string) map[string]int {
	wordCount := make(map[string]int)

	// Split the string into words
	words := strings.Fields(input)

	// Count each word
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func CountWord() {
	str := "go playground and go playground"

	// countWordOne(str)
	// countWordTwo(str)
	countWordsValid(str)

	fmt.Println("")
}
