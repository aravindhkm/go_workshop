package keyword

import "fmt"

type StrType string

func (d *StrType) Update(newVal StrType) {
	*d = newVal
}

type Data struct {
	Name   string
	Mobile int
}

func (d *Data) Update(newStr string, newInt int) {
	d.Name = newStr
	d.Mobile = newInt
}

type IntType int

func (d *IntType) Update(newVal IntType) {
	*d = newVal
}


func PtrOne() {
	strOne := StrType("hello")
	fmt.Println("strOne 1", strOne)

	strOne.Update("Hello world")
	fmt.Println("strOne 2", strOne)

	structOne := Data{"Hello", 1}
	fmt.Println("structOne 1", structOne)

	structOne.Update("Hello World", 2)
	fmt.Println("structOne 2", structOne)

	intOne := IntType(1)
	fmt.Println("intOne 1", intOne)

	intOne.Update(22)
	fmt.Println("intOne 2", intOne)
}

// https://stackoverflow.com/questions/24642311/what-is-the-point-of-passing-a-pointer-to-a-strings-in-go-golang