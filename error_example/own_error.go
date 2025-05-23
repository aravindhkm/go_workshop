package errorexample

import "fmt"

type ErrOwn struct {
	value string
}

func (err ErrOwn) Error() string {
	return err.value
}

func CreateErr(msg string) error {
	return ErrOwn{msg}
}

func isEven(value int) (bool, error) {
	if value%2 == 0 {
		return true, nil
	} else {
		return false, CreateErr("Value is Not Even Value")
	}
}

func OwnErrorMsgOne() {
	result, err := isEven(10)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("result", result)
	}

	result, err = isEven(3)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("result", result)
	}
}
