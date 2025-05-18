package interfacehandle

import "fmt"

type Square struct {
	SideLength float64
}

func (d *Square) getArea() float64 {
	return d.SideLength * d.SideLength
}

type Triangle struct {
	Height float64
	Base   float64
}

func (d *Triangle) getArea() float64 {
	return 0.5 * d.Height * d.Base
}

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {
	fmt.Println("Result", s.getArea())
}
func ItrSix() {
	square := Square{0.5}
	triangle := Triangle{0.5, 4.5}

	printArea(&square)
	printArea(&triangle)
}
