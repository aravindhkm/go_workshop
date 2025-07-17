package interfacehandle

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Text string
}

func (d *Dog) Speak() string {
	d.Text = "Woof!"
	return "Woof!"
}

type Cat struct{}

func (d *Cat) Speak() string {
	return "Meow!"
}

func PrintSound(s Speaker) {
	sound := s.Speak()

	fmt.Println(sound)
}

func InterfaceBasicExample() {
	dog := &Dog{}
	cat := Cat{}

	fmt.Println("dog before", dog.Text)
	PrintSound(dog)
	PrintSound(&cat)

	fmt.Println("dog after", dog.Text)

}
