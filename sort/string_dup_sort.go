package sort

import (
	"fmt"
	"strconv"
	"strings"
)

func StringDuplicateSort() {
	number := 127
	binary := strconv.FormatInt(int64(number), 2)
	paddedBinary := fmt.Sprintf("%08s", binary) // pad with leading zeroes to ensure length of 8
	// fmt.Printf("Binary representation of %d is %s\n", number, paddedBinary)
	count := strings.Count(paddedBinary, "1")
	fmt.Println("count", count)
}
