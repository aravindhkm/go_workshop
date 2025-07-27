package keyword

import "fmt"

func executePanic() {
	fmt.Println("before panic")
	panic("Panic executed successfully")
	fmt.Println("after panic") // it will not execute
}

func arrTry() {
	data := []int{}

	data[4] = 10
}

func accessArray() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[5]) // panic
}

func channelFail() {
	ch := make(chan int)
	close(ch)

	ch <- 10 // panic: send on closed channel
}
// | Scenario                        | Causes Panic? | Can `recover()` help? |
// | ------------------------------- | ------------- | --------------------- |
// | Send on closed channel          | ✅ Yes         | ✅ Yes                 |
// | Receive from closed channel     | ❌ No          | N/A                   |
// | Deadlock (no goroutines active) | ❌ No          | ❌ No                  |
// | Nil channel send/receive        | ✅ Sometimes   | ✅ Yes (with defer)    |

func divide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	fmt.Println(a / b) // panic if b == 0
}

func derefNil() {
	var p *int
	fmt.Println(*p) // panic
}



func recoverExecute() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("Recovered from panic:", msg)
		}
	}()

	// executePanic()
	// arrTry()
	// channelFail()
	// divide(5,0)
	// accessArray()
	derefNil()
}

func RecoverExampleOne() {
	recoverExecute()

	for {
	}
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


// | Scenario                     | Why recover won’t help                  |
// | ---------------------------- | --------------------------------------- |
// | Deadlocks                    | Not a panic, it's a runtime fatal error |
// | Compile-time errors          | Recover only works at runtime           |
// | Memory leaks or logical bugs | Not a panic                             |
// | Infinite loops               | No panic occurs                         |
