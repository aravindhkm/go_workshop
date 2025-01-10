package httphandle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func HttpPostExampleTwo() {
	endpoint := "http://localhost:4000/create"

	payload := map[string]string{
		"Name": "Aravind",
	}

	payloadByte, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("payload error", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadByte))
	if err != nil {
		fmt.Println("request error", err)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("response error", err)
	}

	defer res.Body.Close()

	fmt.Println("res", res)
}