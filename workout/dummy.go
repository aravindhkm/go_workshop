package workout

import (
	"fmt"

	"github.com/google/uuid"
)

func DummyWorkOut() {
	var arr = make([]int, 3)

	arr[0] = 10
	arr[1] = 10
	arr[2] = 10
	// arr[3] = 10

	var slice []string

	for range [4]int{} {
		slice = append(slice, (uuid.NewString()))
	}

	fmt.Println("Arr", arr)
	fmt.Println("Slice", slice)

}