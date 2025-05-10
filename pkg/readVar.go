package pkg

import (
	"fmt"

	"WorkShop/api"
)


func ReadVarPkg() {
	fmt.Println("Reading Variable")
	fmt.Println("Task Variable", api.MyExportedVar)
}