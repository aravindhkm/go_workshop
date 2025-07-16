package interview

import "fmt"

type SafeMap map[int]int

func (d *SafeMap) Add(key, value int) {
	(*d)[key] = value
}

func (d *SafeMap) Get(key int) int {
	if value, ok := (*d)[key]; ok {
		return value
	} else {
		return 0
	}
}

func (d *SafeMap) Delete(key int) bool {
	if _, ok := (*d)[key]; ok {
		delete((*d), key)
		return true
	} else {
		return false
	}
}

func BiMapOne() {
	safeMap := SafeMap{
		5: 10,
	}

	safeMap.Add(4,12)
	safeMap.Add(6,60)

	fmt.Println(safeMap.Get(5))
	fmt.Println(safeMap.Get(4))
	fmt.Println(safeMap.Get(6))

	safeMap.Delete(6)
	fmt.Println(safeMap.Get(6))

	fmt.Println("BiMapOne")
}
