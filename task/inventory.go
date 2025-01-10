//designing a simple inventory management system in Go.
//Design a set of Go types to represent different items in the inventory,
//including Product, Equipment, and Material. Each item should have common attributes like ID,
// Name, and Quantity,
// as well as specific attributes and methods based on their types. Implement methods to update
//quantities, add new items, and list all items in the inventory.

package task

import (
	"fmt"
	"reflect"
)

type Product struct {
	Name string
	Quantity int
}

type Equipment struct {
	Name string
	Quantity int
}

type Material struct {
	Name string
	Quantity int
}

func (a *Product) GetName() string {
	return a.Name
}

func (a *Product) GetQuantity() int {
	return a.Quantity
}

func (a *Product) UpdateQuantity(add int) {
	a.Quantity = add
}


func (a *Equipment) GetName() string {
	return a.Name
}

func (a *Equipment) GetQuantity() int {
	return a.Quantity
}

func (a *Equipment) UpdateQuantity(add int) {
	a.Quantity = add
}

func (a *Material) GetName() string {
	return a.Name
}

func (a *Material) GetQuantity() int {
	return a.Quantity
}

func (a *Material) UpdateQuantity(add int) {
	a.Quantity = add
}

type Item interface {
	GetName() string
	GetQuantity() int
	UpdateQuantity(int)
}

type Inventory struct {
	Items []Item
}

func (a *Inventory) AddItem(item Item) {
	a.Items = append(a.Items, item)
}

func (a *Inventory) ListItems() {
	for _, value := range a.Items {
		fmt.Printf("Product Name: %s\n", value.GetName())
		fmt.Printf("Product Quantity: %v\n", value.GetQuantity())
		fmt.Println("Product DataTypes: ", reflect.TypeOf(value).Name())
	}
}


func InventoryExecute() {
	inventory := Inventory{}

	product_one := Product{Name: "Car", Quantity:100}
	qquipment_one := Equipment{Name: "BackCamera", Quantity:5}
	material_one := Material{Name: "MRF tyre", Quantity:500}

	inventory.AddItem(&product_one)
	inventory.AddItem(&qquipment_one)
	inventory.AddItem(&material_one)

	inventory.ListItems()

	fmt.Println("")
}