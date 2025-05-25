package syncexample

import (
	"fmt"
	"sync"
)

type Kitchen struct {
	maxBufferSize int
	order         []int
}

func NewKitchen(newMaxBufferSize int) *Kitchen {
	return &Kitchen{
		maxBufferSize: newMaxBufferSize,
		order:         make([]int, 0, newMaxBufferSize),
	}
}

func (k *Kitchen) Add(orderId int) {
	k.order = append(k.order, orderId)
}

func (k *Kitchen) Get() int {
	oId := k.order[0]
	k.order = k.order[1:]
	return oId
}

func (k *Kitchen) IsEmpty() bool {
	return len(k.order) == 0
}

func (k *Kitchen) IsFull() bool {
	return len(k.order) == k.maxBufferSize
}

type KProducer struct {
	cond    *sync.Cond
	kitchen *Kitchen
}

func (kp *KProducer) Produce(oId int) {
	kp.cond.L.Lock()
	defer kp.cond.L.Unlock()

	if kp.kitchen.IsFull() {
		fmt.Println("⏸️ -> Producer is waiting because the message channel is full")
		kp.cond.Wait()
	}

	kp.kitchen.Add(oId)
	fmt.Println("➡️  -> Producer produced the order:", oId)

	kp.cond.Signal()
}

type KConsumer struct {
	consumerId int
	cond       *sync.Cond
	kitchen    *Kitchen
}

func (con *KConsumer) Consume() {
	con.cond.L.Lock()
	defer con.cond.L.Unlock()

	if con.kitchen.IsEmpty() {
		fmt.Println("🚫 <Consumer is waiting because the message channel is empty")
		con.cond.Wait()
	}

	orderId := con.kitchen.Get()

	fmt.Printf("⬅️ <Consumer %v consumed the orderId:%v \n", con.consumerId, orderId)
	con.cond.Signal()
}

func Cond_Two_One() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)
	kitchen := NewKitchen(5)

	producer := KProducer{cond, kitchen}

	consumer_one := KConsumer{321, cond, kitchen}
	consumer_two := KConsumer{521, cond, kitchen}

	wg.Add(3)

	go func() {
		defer wg.Done()

		for i := 1; i <= 20; i++ {
			producer.Produce(i)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			consumer_one.Consume()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 11; i <= 20; i++ {
			consumer_two.Consume()
		}
	}()

	wg.Wait()
}
