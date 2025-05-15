package keyword

import "fmt"

type StrType string

func (d *StrType) Update(newVal StrType) {
	*d = newVal
}

func PtrOne() {
	strOne := StrType("hello")
	fmt.Println("strOne 1", strOne)

	strOne.Update("Hello world")
	fmt.Println("strOne 2", strOne)

}
