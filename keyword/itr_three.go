package keyword

import "fmt"

// let check itr_four.go

type Connect interface {
	DbStatus()
	// PodStatus()
	// ServiceStatus()
}

type Db struct {}

func (d *Db) DbStatus() {
	fmt.Println("DB Running Successfully")
}

type Pod struct {}

func (d *Pod) PodStatus() {
	fmt.Println("Pod Running Successfully")
}

type Service struct {}

func (d *Service) ServiceStatus() {
	fmt.Println("Service Running Successfully")
}

func ItrThree() {
	db := &Db{}
	dbItr := Connect(db)
	dbItr.DbStatus()

	// it won't work

	// pod := Pod{}
	// podItr := Connect(pod)
	// podItr.DbStatus()

	// service := Service{}
	// serviceItr := Connect(service)
	// serviceItr.DbStatus()

	fmt.Println("db", db)

}