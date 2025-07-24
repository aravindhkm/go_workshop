package interview

import (
	"fmt"
	"strings"
)

func findAnagramOwn(strOne, strTwo string) bool {
	strOne = strings.ToLower(strOne)
	strTwo = strings.ToLower(strTwo)

	result := false
	for _, one := range strOne {
		// fmt.Println(string(one))
		result = false
		for _, two := range strTwo {
			if two == one {
				result = true
			}
		}
		if !result {
			return false
		}
	}

	return result
}

func findAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s = strings.ToLower(s)
	t = strings.ToLower(t)

	count := [26]int{}

	for i := 0; i < len(s); i++ {
		fmt.Println("", s[i], s[i]-'a')

		count[s[i]-'a']++
		count[t[i]-'a']--

		// count[s[i]-'a']++
		// count[t[i]-'a']--

		// fmt.Println(count[s[i]-'a'], count[t[i]-'a'])
	}

	// If all counts are zero, then it's an anagram
	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}

func AnagramOne() {
	testcase := [][]string{
		{"listen", "silent"},
		// {"gum", "mug"},
		// {"seven", "evens"},
		// {"sdfewfeven", "eveqxns"},
	}

	for _, data := range testcase {
		fmt.Println(data, findAnagram(data[0], data[1]))
	}

}

// Time: O(n), where n is the length of the strings.
// Space: O(1) — because we use a fixed-size array of 26 characters.
