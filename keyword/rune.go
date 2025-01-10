package keyword

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func runeExampleOne() {
	str := "Hello_World_Golang"

	runeStr := []rune(str)
	outputStr := []rune{}

	for i:=len(runeStr)-1; i>=0; i-- {
		outputStr = append(outputStr, runeStr[i])
	}

	fmt.Println("result", string(outputStr))
}

func RuneExecute() {
	runeExampleOne()
	// fmt.Println(reverseString("Hello_World_Golang"))
}