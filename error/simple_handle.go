package error

import (
	"errors"
	"fmt"
)

func isInteger(args interface{}) error {
	fmt.Println("args",args)
	if args != nil {
		fmt.Println("printing")
	}
	if value, status := args.(int); status {
		fmt.Println("value", value)
		fmt.Println("status", status)
		return nil
	} else {
		return errors.New("not a integer")
	}
}


func SimpleErrorHandle() {
	err := isInteger(1)

	if err != nil {
		fmt.Println("error", err)
	}

	err = isInteger("1")
	fmt.Println("before error 1", err)
	if err != nil {
		fmt.Println("error", err)
	}


}