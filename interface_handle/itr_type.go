package interfacehandle

import "fmt"

func findItr(itrData interface{}) {
	switch v := itrData.(type) {
	case int:
		fmt.Println("value is integer", v)
	case string:
		fmt.Println("value is string", v)
	case bool:
		fmt.Println("value is bool", v)
	default:
		fmt.Println("value is of unknown type", v)
	}
}

func ItrTypeFind() {
	data := 5
	findItr(data)

	dataStr := "hello"
	findItr(dataStr)

	var dataItr bool
	findItr(dataItr)
}
