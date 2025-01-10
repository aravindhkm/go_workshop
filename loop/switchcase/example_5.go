package switchcase

import "fmt"

type Data struct {
	A      int
	B      int
	Result int
}

func (d *Data) add(a int, b int) {
	d.A = a
	d.B = b
	d.Result = d.A + d.B
}

func (d *Data) operator() {
	d.add(d.A, d.B)
}

type SubData struct {
	A      int
	B      int
	Result int
}

func (d *SubData) sub(a int, b int) {
	d.A = a
	d.B = b
	d.Result = d.A - d.B
}

func (d *SubData) operator() {
	d.sub(d.A, d.B)
}

type operator interface {
	operator()
}

func ExampleFive() {
	addition_one := &Data{A: 4, B: 6}
	addition_two := &SubData{A: 4, B: 6}

	op_itf := []operator{addition_one, addition_two}

	for _, value := range op_itf {
		// fmt.Println("index", index)

		switch op := value.(type) {
		case *Data:
			{
				op.add(op.A, op.B)
				fmt.Printf("Addition: %d + %d = %d\n", op.A, op.B, op.Result)
			}
		case *SubData:
			{
				op.sub(op.A, op.B)
				fmt.Printf("Subtraction: %d - %d = %d\n", op.A, op.B, op.Result)
			}
		}
	}
}
