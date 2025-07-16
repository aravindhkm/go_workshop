package generics

import "fmt"

type AnyMap[K comparable, V comparable] map[K]V

type AnyType[T comparable] string

func ComparableTwo() {
	intMap := AnyMap[string, int]{
		"A": 4,
		"B": 2,
	}

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	strMap := AnyMap[int, bool]{
		111: true,
		5:   false,
	}

	for key, value := range strMap {
		fmt.Println(key, value)
	}

	var name AnyType[int] = "Hello"
	fmt.Println(name)

	var another AnyType[string] = "World"
	fmt.Println(another)
}
