package loop

import (
	"WorkShop/loop/selectcase"
	"WorkShop/loop/switchcase"
	"fmt"
)

const context = "switchcase"

func Run(funcName string) {

	if context == "selectcase" {
		selectcase.SelectCaseExecute()
	}

	if context == "switchcase" {
		switchcase.SwitchExecute()
	}

	fmt.Println("")
}
