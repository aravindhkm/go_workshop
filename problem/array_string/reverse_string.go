/*
Problem:
- Given a list of string, reverse its order in place.

Example:
- Input: []string{"a", "b", "c", "d"}
  Output: []string{"d", "c", "b", "a"}

Approach:
- Use two pointers approach to swap two values on both ends as we move toward
  the middle.

Solution:
- Initialize the two pointers, one starts at the beginning and one starts at
  the end of the list.
- While the start pointer does not meet the end pointer in the middle, keep
  swapping these two values.
- Move the start pointer up and move the end pointer down.

Cost:
- O(n) time, O(1) space.
*/

package arraystring

import "fmt"

func MyReserveString() {
	input := []string{"a", "b", "c", "d"}
	output := []string{}

	for i:=len(input)-1; i>=0; i-- {
		output = append(output, input[i])
	}

	fmt.Println("output", output)

	test := []struct{
		in []string
		out []string
	}{
		{
			[]string{"a","b"},
			[]string{"c","d"},
		},
		{
			[]string{"e","f"},
			[]string{"g","h"},
		},
	}

	fmt.Println("test", test)
}

func actualResult() {
	input := []string{"a", "b", "c", "d"}

	start := 0
	end := len(input) - 1

	for start < end {
		input[start], input[end] = input[end], input[start]

		start++
		end--
	}
	fmt.Println("input", input)
}

func ReverseString() {
	// MyReserveString()
	actualResult()
}