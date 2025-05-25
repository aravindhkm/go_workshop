package keyword

import "fmt"

func Copy_B() {
	src := []int{1, 2, 3, 4, 5}
	dst := src[1:4]
	fmt.Println("before dst", dst, src[2:])
	// dst -> 2,3,5 -> 1,2,3
	// src -> 3,4,5 -> 2,4,5
	copy(dst, src[2:])
	fmt.Println("src:", src) // 1,3,4,5,5
	fmt.Println("dst:", dst) // 3,4,5
}


// ✅ Important: dst still refers to the same underlying array as src.

// Copy Step  From src[2:]	        To dst	             Updated src
// i = 0	   src[2] = 3	  dst[0] → src[1] = 3	[1, 3, 3, 4, 5]
// i = 1	   src[3] = 4	  dst[1] → src[2] = 4	[1, 3, 4, 4, 5]
// i = 2	   src[4] = 5	  dst[2] → src[3] = 5	[1, 3, 4, 5, 5]
