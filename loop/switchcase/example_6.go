package switchcase

import (
	"errors"
	"fmt"
)

const (
	One = iota
	Two
)

type Calculator interface {
	Add(args ...int) (int, error)
	Sub(args ...int) (int, error)
}

type Value struct {
	A int
	B int
}

type Value1 struct {
	A int
	B int
	C int
	Value
}

func calculate(c Calculator, method int) (int, error) {
	switch method {
	case One:
		if v1, ok := c.(Value1); ok {
			return c.Add(v1.A, v1.B, v1.C)
		}
		sum, err := c.Add()
		if err != nil {
			return 0, err
		}
		return sum, nil
	case Two:
		if v1, ok := c.(Value1); ok {
			return c.Sub(v1.A, v1.B, v1.C)
		}
		diff, err := c.Sub()
		if err != nil {
			return 0, err
		}
		if v1, ok := c.(Value1); ok {
			return diff - v1.C, nil
		}
		return diff, nil
	default:
		return 0, fmt.Errorf("invalid method: %v", method)
	}
}

func (v Value) Add(args ...int) (val int, err error) {
	if len(args) == 0 {
		if v.A == 0 || v.B == 0 {
			return 0, errors.New("passed 0 as an arg")
		}
		return v.A + v.B, nil
	}
	for _, v := range args {
		val += v
	}
	return val, nil
}

func (v Value) Sub(args ...int) (val int, err error) {
	if len(args) == 0 {
		if v.A == 0 || v.B == 0 || v.A <= v.B {
			return 0, errors.New("invalid args")
		}
		return v.A - v.B, nil
	}
	// [3, 1, 2]
	val = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > val {
			return 0, errors.New("underflow")
		}
		val -= args[i]
	}
	return val, nil
}

func ExampleSix() {
	v := Value{A: 5, B: 4}
	res, err := calculate(v, One)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("add value :", res)
	res, err = calculate(v, Two)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sub value: ", res)
	fmt.Println()
	v1 := Value1{A: 3, B: 1, C: 2}
	res, err = calculate(v1, One)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v1 add: ", res)
	res, err = calculate(v1, Two)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v1 sub: ", res)
	fmt.Println()
	fmt.Println("backshot over to Siva...")
}
