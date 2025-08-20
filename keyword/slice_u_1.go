package keyword

import (
	"fmt"
)

func SliceUnderLyingOne() {
	// Create a slice with values 0 to 9.
	// Backed by an underlying array of size 10.
	slice := []int{0, 1, 2, 3, 4, 5,6,7,8}
	fmt.Println("Initial slice:", slice, "len:", len(slice), "cap:", cap(slice))

	// Take a subslice of the first 5 elements (indexes 0..4).
	// NOTE: subSlice shares the same underlying array as 'slice'.
	// So, changes through subSlice may also reflect in 'slice'.
	subSlice := slice[:5]
	fmt.Println("Initial subSlice:", subSlice, "len:", len(subSlice), "cap:", cap(subSlice))

	// Append 500 to subSlice.
	// Since subSlice still has free capacity (cap=10, len=6),
	// the new element is placed in the same backing array,
	// modifying slice[5] = 500.
	subSlice = append(subSlice, 500)
	fmt.Println("After append 500 -> subSlice:", subSlice, "len:", len(subSlice), "cap:", cap(subSlice))
	fmt.Println("Slice after append 500:", slice, "len:", len(slice), "cap:", cap(slice))

	// Append 600 to subSlice.
	// Again, within capacity, so it modifies slice[6] = 600.
	subSlice = append(subSlice, 600)
	fmt.Println("After append 600 -> subSlice:", subSlice, "len:", len(subSlice), "cap:", cap(subSlice))
	fmt.Println("Slice after append 600:", slice, "len:", len(slice), "cap:", cap(slice))

	// Append 700 to subSlice.
	// Still within the same underlying array, so slice[7] = 700 changes.
	subSlice = append(subSlice, 700)
	fmt.Println("After append 700 -> subSlice:", subSlice, "len:", len(subSlice), "cap:", cap(subSlice))
	fmt.Println("Slice after append 700:", slice, "len:", len(slice), "cap:", cap(slice))

	// Append 800 to subSlice.
	// Still fits, so slice[8] = 800 changes.
	subSlice = append(subSlice, 800)
	fmt.Println("After append 800 -> subSlice:", subSlice, "len:", len(subSlice), "cap:", cap(subSlice))
	fmt.Println("Slice after append 800:", slice, "len:", len(slice), "cap:", cap(slice))

	// Append 900 to subSlice.
	// NOW subSlice length would exceed the capacity of the shared array.
	// Go creates a NEW underlying array (capacity usually doubled to 20).
	// subSlice now points to this new array and stops sharing memory with 'slice'.
	// So, further changes to subSlice do not affect slice.
	subSlice = append(subSlice, 900)
	fmt.Println("After append 900 -> subSlice:", subSlice, "len:", len(subSlice), "cap:", cap(subSlice))
	fmt.Println("Slice after append 900:", slice, "len:", len(slice), "cap:", cap(slice))

	fmt.Println("")
	fmt.Println(slice)
	fmt.Println(subSlice)
}