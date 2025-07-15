package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type WriteJsonStruct struct {
	Data    int    `json:"data"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func WriteJsonExampleOne() {

	op := WriteJsonStruct{200, "Successfully writed", "Hello Aravindh!"}

	data, err := json.Marshal(op)

	if err != nil {
		fmt.Println("err", err)
	}
	perm := os.FileMode(0644) // This represents -rw-r--r--

	os.WriteFile("output.json", data, perm)

	fmt.Println("op", op)

}

//
// 0 means no permissions.
// 1 means execute permission.
// 2 means write permission.
// 3 means write and execute permissions.
// 4 means read permission.
// 5 means read and execute permissions.
// 6 means read and write permissions.
// 7 means read, write, and execute permissions.

// It's a Unix-style file permission:

// 0: special indicator for octal

// 6: owner can read and write (4+2)

// 4: group can read

// 4: others can read
