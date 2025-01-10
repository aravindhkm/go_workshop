package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type ResultJsonTwo struct {
	Data int `json:"data"`
	Message string `json:"message"`
	Result string `json:"result"`
}

func ReadJsonFileTwo(){
	var result ResultJsonTwo
	data, err := os.Open("output.json")
	
	if err != nil {
		fmt.Println("Err :", err)
	}

	err = json.NewDecoder(data).Decode(&result)

	if err != nil {
		fmt.Println("unmarshal err :", err)
	}

	fmt.Println(result)
}