package keyword

import "fmt"

func DeferExecute() {
	defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")

	fmt.Println("Defer start")
}