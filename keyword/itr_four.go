package keyword

import "fmt"

type ConnectItr interface {
	Status()
}

type DbItr struct{}

func (d *DbItr) Status() {
	fmt.Println("DB Running Successfully")
}

type PodItr struct{}

func (d *PodItr) Status() {
	fmt.Println("Pod Running Successfully")
}

type ServiceItr struct{}

func (d *ServiceItr) Status() {
	fmt.Println("Service Running Successfully")
}

func ItrFour() {
	dbItrOne := &DbItr{}
	podItrOne := &PodItr{}
	serviceItrOne := &ServiceItr{}

	dbConnectItr := ConnectItr(dbItrOne)
	dbConnectItr.Status()

	podConnectItr := ConnectItr(podItrOne)
	podConnectItr.Status()

	serviceConnectItr := ConnectItr(serviceItrOne)
	serviceConnectItr.Status()
}
