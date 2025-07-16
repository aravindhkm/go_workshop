package interview

import "fmt"

func facGet(n int) int {
	for n == 0 || n == 1 {
		return 1
	}
	return n * facGet(n-1)
}

func FacOne() {
	fmt.Println("Factorial", facGet(5))
}


// | Concept              | Explanation                                                                   |
// | -------------------- | ----------------------------------------------------------------------------- |
// | **Recursion**        | The function calls itself with `n-1` until it hits base case `n == 0` or `1`. |
// | **Call Stack**       | Each recursive call is pushed onto the system’s call stack.                   |
// | **Base Case**        | Stops recursion (when `n == 0` or `n == 1`).                                  |
// | **Time Complexity**  | O(n) — one recursive call for each `n` down to 1.                             |
// | **Space Complexity** | O(n) — due to call stack space (not efficient for very large `n`).            |
