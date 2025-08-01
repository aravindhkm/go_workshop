package keyword

import "fmt"

func appendGrow() {
	arr := make([]int, 5)
	fmt.Println("a", arr, len(arr), cap(arr))
	arr = append(arr, 2, 5)
	fmt.Println("b", arr, len(arr), cap(arr))

	arr = append(arr, 7)
	fmt.Println("c", arr, len(arr), cap(arr))

	arr = append(arr, 8, 9)
	fmt.Println("d", arr, len(arr), cap(arr))

	arr = append(arr, 10)
	fmt.Println("e", arr, len(arr), cap(arr))
}

func loopGrow() {
	s := make([]int, 0)
	oldCap := cap(s)

	for i := 0; i < 2000; i++ {
		s = append(s, i)
		if cap(s) != oldCap {
			fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
			oldCap = cap(s)
		}
	}

	// len=1 cap=1
	// len=2 cap=2
	// len=3 cap=4
	// len=5 cap=8
	// len=9 cap=16
	// len=17 cap=32
	// len=33 cap=64
	// len=65 cap=128
	// len=129 cap=256
	// len=257 cap=512
	// len=513 cap=848
	// len=849 cap=1280
	// len=1281 cap=1792
	// len=1793 cap=2560
	// 	Note the doubling up to 1024.
	// Then it switches to 1.25x growth.

	// | Capacity Size | Growth Pattern     |
	// | ------------- | ------------------ |
	// | `< 1024`      | Doubles (×2)       |
	// | `≥ 1024`      | 1.25× (\~25% more) |

	// Efficiency: Doubling for small slices keeps allocations low.
	// Memory Safety: 1.25x growth for large slices prevents wasting too much memory.

}

func CapOne() {
	// appendGrow()
	loopGrow()
}
