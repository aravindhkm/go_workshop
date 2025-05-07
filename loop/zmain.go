package loop

import (
	"WorkShop/loop/selectcase"
	"WorkShop/loop/switchcase"
)

var floderName = map[string]func(string){
	"selectcase": selectcase.Run,
	"switchcase": switchcase.Run,
}

var execFloder string = "switchcase"

// Run executes a function by key from the store
func Run(funcName string) {
	if fn, exists := floderName[execFloder]; exists {
		fn(funcName)
		return
	} else {
		floderName[execFloder](funcName)
		return
	}
}
