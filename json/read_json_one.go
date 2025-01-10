package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type ResultJson struct {
	Status int
	Data string
}

func ReadJsonFile(){
	var result ResultJson
	data, err := os.ReadFile("data.json")
	
	if err != nil {
		fmt.Println("Err :", err)
	}

	err = json.Unmarshal(data,&result)

	if err != nil {
		fmt.Println("unmarshal err :", err)
	}

	fmt.Println(result.Data, result.Status, result)
}