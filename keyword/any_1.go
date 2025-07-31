package keyword

import "fmt"

// Using 'any'
func PrintValue[T any](val T) {
	fmt.Println(val)
}

// Using 'any'
func PrintItf[T interface{}](val T) {
	fmt.Println(val)
}

// Generic filter function using `any`
func FilterAny[T any](list []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func FilterAnyExec() {
	numbers := []int{1, 2, 3, 4, 5}
	// Keep only even numbers
	evens := FilterAny(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evens)
}

// Same filter function, using `interface{}` instead of `any`
func FilterItf[T interface{}](list []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func FilterItfExec() {
	numbers := []int{1, 2, 3, 4, 5}
	evens := FilterItf(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evens)
}

func AnyOne() {
	// PrintValue(42)          // int
	// PrintValue("Hello")     // string
	// PrintValue([]int{1, 2}) // slice

	// FilterAnyExec()
	FilterItfExec()
}
