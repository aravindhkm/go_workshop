package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type WriteJsonStructTwo struct {
	Data int `json:"data"`
	Message string `json:"message"`
	Result string `json:"result"`
}

func WriteJsonExampleTwo() {

	op := WriteJsonStructTwo{200,"Successfully writed", "Hello Aravindh!"}

	file, err := os.Create("output.json")

	if err != nil {
		fmt.Println("Write err", err)
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(&op)

	if err != nil {
		fmt.Println("Write Data err", err)
	}
	fmt.Println("op",op)

}
