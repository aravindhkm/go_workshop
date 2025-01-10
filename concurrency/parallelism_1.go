package concurrency


import (
    "fmt"
)

func printNumbers(from, to int, ch chan int) {
    for i := from; i <= to; i++ {
        ch <- i // Send 'i' to the channel
    }
}

func ParallelismOne() {
    // Create a channel of integers
    ch := make(chan int)

    // Launch two goroutines to print numbers in parallel
    go printNumbers(1, 5, ch)
    go printNumbers(6, 10, ch)

    // Read from the channel
    for i := 0; i < 10; i++ {
        fmt.Println(<-ch) // Receive data from the channel
    }
}
