package interview

import (
	"fmt"
)

func findDupChar(itr []interface{}) []string {
	result := []string{}

	// for _, value := range itr {
	// 	switch value.(Type) {
	// 	case string:
	// 	}
	// 	fmt.Println(reflect.TypeOf(value))
	// }

	return result
}

func FindCharDuplicate() {
	nameList := []interface{}{ // Output - [jhon ken steve jared tom]
		"jhon",
		"ken",
		"steve",
		map[string]string{"name": "jared"},
		"Jhon",
		map[string]string{"name": "steve"},
		"ken",
		"tom",
	}

	result := findDupChar(nameList)

	fmt.Println(result)
}
