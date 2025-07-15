package main

import (
	"WorkShop/config"
	"WorkShop/utils"
	"log"
	"strings"

	"fmt"
	"os"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <your_string>")
		return
	}

	pkgName := os.Args[1]
	funcName := ""

	if len(os.Args) >= 3 {
		funcName = os.Args[2]
		if pkgName != "zexecute/" && strings.Contains(funcName, "l") {
			config.SetReverse(true)
			funcName = strings.ReplaceAll(funcName, "l", "")
		}
	} else {
		funcName = "0"
	}

	execFunc := utils.FindRunMap(pkgName)
	execFunc(funcName)
}
