package task

import "fmt"

func fibTwo(a, b, loop int) int {
	if loop < 10 {
		loop++
		fmt.Println(a)
		return fibTwo(a+b, a, loop)
	}
	return a
}

func fibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	return memo[n]
}

func fibDsa() {
	n := 10 // Change n to get more Fibonacci numbers
	memo := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibonacci(i, memo))
	}
}

func FibonacciTwo() {
	// a := fibTwo(0,1, 0)
	// fmt.Println(a)
	// fmt.Println("FibonacciTwo")

	fibDsa()
}
