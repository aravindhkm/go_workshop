package httphandle

import "fmt"

func IoReaderOne() {
	b := make([]byte, 0, 512)
	fmt.Println("len(b):cap(b)", len(b), cap(b))
}