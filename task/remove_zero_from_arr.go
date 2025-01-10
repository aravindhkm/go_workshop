package task

import "fmt"


func removeEleFromArr() {
	data := []int{1,0,4,6,7,0,0,4,3,0,5,0}
	pos := 0

	for _, value := range data {
		if value != 0 {
			data[pos] = value
			pos++
		}
	}

	for pos < len(data) {
		data[pos] = 0
		pos++
	}

	fmt.Println(data)
}

func RemoveZeroFromArray() {
	removeEleFromArr()
}