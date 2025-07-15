package interview

import (
	"fmt"
	"strings"
	// "reflect"
)

func findDupChar(itr []interface{}) []string {
	result := []string{}
	mapKey := map[string]bool{}
	for _, value := range itr {
		switch c := value.(type) {
		case string:
			mapKey[strings.ToLower(c)] = true
		case map[string]string:
			for _, val := range c {
				mapKey[val] = true
			}
		}
	}

	for key := range mapKey {
		result = append(result, key)
	}

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
