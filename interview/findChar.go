package interview

import "fmt"

func FindChar() {
	str := "Hello3,  World!!123"

	var vowels, spaces, specials, consonants int

	for _, ch := range str {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			switch ch {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				vowels++
			default:
				consonants++
			}
		} else if ch == ' ' {
			spaces++
		} else {
			specials++
		}
	}

	fmt.Println("result", vowels, spaces, specials, consonants)
}
