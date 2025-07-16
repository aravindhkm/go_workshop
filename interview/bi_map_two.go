package interview

import "fmt"

type SafeCompMap[K comparable, V comparable] map[K]V

func (d *SafeCompMap[K, V]) Add(key K, value V) {
	(*d)[key] = value
}

func (d *SafeCompMap[K, V]) Get(key K) V {
	value, _ := (*d)[key]
	return value
}

func (d *SafeCompMap[K, V]) Delete(key K) bool {
	if _, ok := (*d)[key]; ok {
		delete((*d), key)
		return true
	} else {
		return false
	}
}
func BiMapTwo() {
	safeMap := SafeCompMap[int, int]{
		5: 10,
	}

	safeMapTwo := SafeCompMap[string, bool]{
		"hello": true,
	}

	safeMap.Add(4, 12)
	safeMap.Add(6, 60)

	safeMapTwo.Add("Test", true)

	fmt.Println(safeMap.Get(5))
	fmt.Println(safeMap.Get(4))
	fmt.Println(safeMap.Get(6))

	safeMap.Delete(6)
	fmt.Println(safeMap.Get(6))

	fmt.Println("BiMapOne")
}
