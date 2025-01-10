package loop

import (
	"fmt"
    "WorkShop/loop/switchcase"
	"WorkShop/loop/selectcase"

)

const context = "switchcase"

func LoopExecute() {
	
	if context == "selectcase"{
		selectcase.SelectCaseExecute()
	}

	if context == "switchcase"{
		switchcase.SwitchExecute()
	}
	
	fmt.Println("")
}