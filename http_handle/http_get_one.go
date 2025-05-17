package httphandle

import (
	"encoding/json"
	"fmt"
	"io"

	// "log"
	"net/http"
)

func HttpGetOne() {
	url := "https://restcountries.com/v3.1/name/India"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("Status", resp.StatusCode)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// log.Fatalf("Failed to read response body: %v", err)
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		// log.Fatalf("Failed to parse JSON: %v", err)
	}

	// fmt.Println("Random Dog Image URL:", result)

	for v := range 10 {
		fmt.Println("print", v)
	}
}
