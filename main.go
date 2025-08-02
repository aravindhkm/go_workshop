package main

import (
	"WorkShop/config"
	"WorkShop/utils"
	"flag"
	"log"
	"strings"

	"fmt"
	"os"
)

func main() {
	// Parse command-line flag for file path input
	filePathArg := flag.String("args", "task/", "Specify the input file directory")
	fileIndexArg := flag.String("index", "0", "Specify the file index")
	flag.Parse()

	// fmt.Println("Application started with args:", *filePathArg)

	// Split and display parts of the file path
	pkgPath := strings.Split(*filePathArg, "/")
	// fmt.Println("Parsed file path segments:", pkgPath, pkgPath[0], pkgPath[1])

	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Configuration loading failed: %v", err)
	}

	// Check command-line arguments for execution context
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <packageName> [functionName]")
		os.Exit(1)
	}
	pkgName := pkgPath[0]
	funcName := ""

	if *fileIndexArg != "0" && *fileIndexArg != "" {
		config.SetIndexFilter(true)

		if strings.Contains(*fileIndexArg, "l") {
			config.SetReverse(true)
			funcName = strings.ReplaceAll(*fileIndexArg, "l", "")
		} else {
			funcName = *fileIndexArg
		}
	} else if len(pkgPath) > 1 {
		funcName = pkgPath[1]
	} else {
		log.Fatalf("Invalid File Index")
	}

	// Retrieve and invoke the corresponding execution function
	execFunc := utils.FindRunMap(pkgName)
	execFunc(funcName)
}
