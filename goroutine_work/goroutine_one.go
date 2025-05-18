package goroutinework

import (
	"fmt"
	"net/http"
)

func MakeCall(urls []string) {
	for _, endpoint := range urls {
		res, err := http.Get(endpoint)
		if err != nil {
			fmt.Printf("Error fetching URL %s: %v\n", endpoint, err)
		}

		fmt.Println(endpoint, res.StatusCode)
	}
}
func GoRoutineOne() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://api.github.com", // public GitHub API
		"https://httpbin.org/get",
	}

	MakeCall(urls)
}
