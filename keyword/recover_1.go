package keyword

import "fmt"


func executePanic() {
	fmt.Println("before panic")
	panic("Panic executed successfully")
	fmt.Println("after panic") // it will not execute
}

func recoverExecute() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("Recovered from panic:", msg)
		}
	}()
	executePanic()
}

func RecoverExampleOne() {
	recoverExecute()

	for{}
}

// negative scenario
// func recoverExecute() {
// 	defer func() {
// 		if msg := recover(); msg != nil {
// 			fmt.Println("Recovered from panic:", msg)
// 		}
// 	}()
// }

// func RecoverExampleOne() {
// 	recoverExecute()

// 	panic("error main")
// }