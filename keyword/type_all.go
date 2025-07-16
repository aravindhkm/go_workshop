package keyword

import "fmt"

type IntSafe int
type StrSafe string
type BoolSafe bool
type ItrSafe interface{}

type SingleIntArr []int
type DoubleIntArr [][]int

type SingleStringArr []string
type DoubleStringArr [][]string

type MapSafe map[int]int
type MapSafeTwo map[string]bool

func TypeAll() {
	safeInt := IntSafe(4)
	fmt.Println("Int", safeInt)

	safeStr := StrSafe("Hello")
	fmt.Println("STRING", safeStr)

	safeBool := BoolSafe(true)
	fmt.Println("BOOL", safeBool)

	safeItr := ItrSafe("Interface Test")
	fmt.Println("INTERFACE", safeItr)

	safeIntArr := SingleIntArr{4, 1, 35, 6}
	fmt.Println("SingleIntArr", safeIntArr)

	safeDobArr := DoubleIntArr{
		{4, 1},
		{35, 6},
	}
	fmt.Println("DoubleIntArr", safeDobArr)

	safeStrArr := SingleStringArr{"fwefe", "efe", "hub"}
	fmt.Println("SingleStringArr", safeStrArr)

	safeDobStrArr := DoubleStringArr{
		{"4, 1", "fefre"},
		{"test", "hello", "world"},
	}
	fmt.Println("DoubleStringArr", safeDobStrArr)

	safeMap := MapSafe{
		1: 5,
	}
	fmt.Println("SafeMap", safeMap)

	safeMapTwo := MapSafeTwo{
		"hello": true,
	}
	fmt.Println("SafeMapTwo", safeMapTwo)
	// fmt.Println("TypeAll")
}
