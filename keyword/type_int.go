package keyword

import "fmt"

type IntSlice []int

func (d *IntSlice) AddSafe(newVal int) {
	*d = append(*d, newVal)
}

func (d *IntSlice) RemoveIndexSafe(index int) bool {
	if index >= 0 && index < len(*d) {
		(*d) = (*d)[:len((*d))-index]
		return true
	} else {
		return false
	}
}

func (d *IntSlice) RemoveValueSafe(value int) {
	i := 0
	for _, data := range *d {
		if data != value {
			(*d)[i] = data
			i++
		}
	}
	*d = (*d)[:i]
}

func (d *IntSlice) GetByIndex(index int) (int, bool) {
	if index >= 0 && index < len(*d) {
		return (*d)[index], true
	} else {
		return (*d)[0], false
	}
}

func (d *IntSlice) Print() {
	for index, value := range *d {
		fmt.Println("Index :", index, "Value :", value)
	}
}

func (d *IntSlice) Length() int {
	return len((*d))
}

func removeByValueInPlace(slice []int, value int) []int {
	i := 0
	for _, v := range slice {
		if v != value {
			slice[i] = v
			i++
		}
	}
	return slice[:i]
}

func TypeInt() {
	data := &IntSlice{}
	data.AddSafe(1)
	data.AddSafe(5)
	data.AddSafe(7)

	val, ok := data.GetByIndex(2)
	if ok {
		fmt.Println("Value:", val)
	} else {
		fmt.Println("Invalid index")
	}
	data.Print()
	fmt.Println("Length", data.Length())

	// isRemoved := data.RemoveIndexSafe(2)
	// fmt.Println("isRemoved", isRemoved)

	data.RemoveValueSafe(5)

	data.Print()
	fmt.Println("Length", data.Length())

	// sliceVal := []int{2, 3}

	// fmt.Println("length slice 1", sliceVal, len(sliceVal))

	// sliceVal = append(sliceVal, 4, 7, 8)

	// fmt.Println("length slice 2", sliceVal, len(sliceVal))

	// sliceVal = sliceVal[:len(sliceVal)-1]
	// fmt.Println("length slice 3", sliceVal, len(sliceVal))

	// sliceVal = append(sliceVal, 10)

	// fmt.Println("length slice 4", sliceVal, len(sliceVal))

	// // sliceVal = sliceVal[2:5]

	// sliceVal = removeByValueInPlace(sliceVal, 4)

	// fmt.Println("length slice 5", sliceVal, len(sliceVal))

}
