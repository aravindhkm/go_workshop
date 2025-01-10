/*
Problem:
- Given a list of string that is made up of word but in reverse, return the
  correct order in-place.

Example:
- Input: []string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}
  Output: []string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"}

Approach:
- Similar to reversing string, use the same idea to reverse all the characters
  in the list first so that we could have a list of words in the right order, not
  its characters.
- Iterate through the list again and reverse its characters.

Solution:
- Reverse all the characters to get a list of words in the right order using
  same idea as reversing a string.
- Iterate through the list again the reverse its characters by first keeping
  track of the start index for each word since they are separated by an empty
  string.
- Once we fine an empty string, update the start word index and use the same
  idea to reverse the characters order.

Cost:
- O(n) time, O(1) space.
*/

package arraystring

import "fmt"


func expectedReverseWords() {
	arr := []string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}

	// Helper function to reverse a section of the array in place
	reverse := func(arr []string, start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	// Step 1: Reverse the entire array to get the words in the right order
	reverse(arr, 0, len(arr)-1)
	fmt.Println("arr 1", arr)

	// Step 2: Reverse each word individually
	start := 0
	for i := 0; i <= len(arr); i++ {
		// When we find an empty string or reach the end of the array,
		// reverse the word that ends right before this point
		if i == len(arr) || arr[i] == "" {
			reverse(arr, start, i-1)
			// Update start to the next character after the empty string
			start = i + 1
		}
	}

	fmt.Println("arr", arr)
}


func MyReverseWord() {
	input := []string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}
	findIndex := []int{0}

	for index, value := range input {
		if value == "" {
			findIndex = append(findIndex, index+1)
		}
	}

	result := []string{}

	for i:= len(findIndex)-1; i>=0; i-- {
		start := findIndex[i]

		for j:=start; j<len(input); j++ {
			result = append(result, input[j])

			if input[j] == "" {
				break
			}

			if j == len(input) -1 {
				result = append(result, "")

			}
		}

	}

	fmt.Println(result)
}

func ReverseWord() {
	// MyReverseWord()
	expectedReverseWords()
}
