package httphandle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func HttpPostExampleOne() {
	url := "https://dummy.restapiexample.com/api/v1/create"

	reqBody := map[string]string{
		"data": "create",
	}

	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	fmt.Println("Response status:", response.StatusCode)

	var apiResult map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&apiResult); err != nil {
		fmt.Println("Error decoding response:", err.Error())
		return
	}

	fmt.Println("API Result:", apiResult)
}
