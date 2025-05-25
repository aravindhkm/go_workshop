package syncexample

import (
	"fmt"
	"sync"
)

func printMap(key, value any) bool {
	fmt.Println("Key --->", key, " Value --->", value)
	return true
}

func NormalSafeMapTest() {
	safeMap := sync.Map{}

	safeMap.Store("Ak", "Hello")
	if value, okay := safeMap.Load("Ak"); okay {
		fmt.Println("Ak:", value)
	}

	safeMap.Store(1, 10)
	if value, okay := safeMap.Load(1); okay {
		fmt.Println("1:", value)
	}

	safeMap.Store(true, 4.5)
	if value, okay := safeMap.Load(true); okay {
		fmt.Println("true:", value)
	}

	safeMap.Range(printMap)

	safeMap.Delete(true)

	safeMap.Range(printMap)

	safeMap.Clear()
	fmt.Println("PrintRange")
	safeMap.Range(printMap)

}

func Map_B() {
	NormalSafeMapTest()
}
