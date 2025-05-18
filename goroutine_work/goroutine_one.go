package goroutinework

import (
	"fmt"
	"net/http"
)

type Result struct {
	endpoint   string
	statusCode int
}

func MakeCall(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
	}

	// fmt.Println(url, res.StatusCode)
	tmpRes := Result{url, res.StatusCode}
	ch <- tmpRes

}
func GoRoutineOne() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://api.github.com", // public GitHub API
		"https://httpbin.org/get",
	}

	ch := make(chan Result)

	for _, endpoint := range urls {
		go MakeCall(endpoint, ch)
	}

	// channel receiver is a blocking area

	fmt.Println("end result 1", <-ch)


	// fmt.Println("end result 2", <-ch)
	// fmt.Println("end result 3", <-ch)
	// fmt.Println("end result 3", <-ch)

}
