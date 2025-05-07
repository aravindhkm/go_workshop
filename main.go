package main

import (
	"WorkShop/utils"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <your_string>")
		return
	}

	pkgName := os.Args[1]
	funcName := ""
	
	if len(os.Args) >= 3 {
		funcName = os.Args[2]
	}

	execFunc := utils.FindRunMap(pkgName)
	execFunc(funcName)
}
