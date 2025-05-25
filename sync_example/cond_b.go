package syncexample

import (
	"fmt"
	"sync"
	"time"
)

type MessageChannel struct {
	maxBufferSize int
	buffer        []string
}

func NewMessageChannel(size int) *MessageChannel {
	return &MessageChannel{
		maxBufferSize: size,
		buffer:        make([]string, 0, size),
	}
}

func (mc *MessageChannel) IsEmpty() bool {
	return len(mc.buffer) == 0
}

func (mc *MessageChannel) IsFull() bool {
	return len(mc.buffer) == mc.maxBufferSize
}

func (mc *MessageChannel) Add(message string) {
	mc.buffer = append(mc.buffer, message)
}

func (mc *MessageChannel) Get() string {
	message := mc.buffer[0]
	mc.buffer = mc.buffer[1:]
	return message
}

type Producer struct {
	cond           *sync.Cond
	messageChannel *MessageChannel
}

func NewProducer(cond *sync.Cond, messageChannel *MessageChannel) *Producer {
	return &Producer{
		cond:           cond,
		messageChannel: messageChannel,
	}
}

func (p *Producer) Produce(message string) {
	time.Sleep(500 * time.Millisecond) // Simulating some work

	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	for p.messageChannel.IsFull() {
		fmt.Println("🚫 -> Producer is waiting because the message channel is full")
		p.cond.Wait()
	}

	p.messageChannel.Add(message)
	fmt.Println("➡️  -> Producer produced the message:", message)

	p.cond.Signal()
}

type Consumer struct {
	id             int
	cond           *sync.Cond
	messageChannel *MessageChannel
}

func NewConsumer(cId int, cond *sync.Cond, messageChannel *MessageChannel) *Consumer {
	return &Consumer{
		id:             cId,
		cond:           cond,
		messageChannel: messageChannel,
	}
}

func (c *Consumer) Consume() {
	time.Sleep(1 * time.Second) // Simulating some work

	c.cond.L.Lock()
	defer c.cond.L.Unlock()

	for c.messageChannel.IsEmpty() {
		fmt.Println("⏸️ <Consumer is waiting because the message channel is empty")
		c.cond.Wait()
	}

	message := c.messageChannel.Get()

	fmt.Println("⬅️ <Consumer consumed the message:", message)
	c.cond.Signal()
}

const MaxMessageChannelSize = 5

func Cond_Two() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	messageChannel := NewMessageChannel(MaxMessageChannelSize)
	producer := NewProducer(cond, messageChannel)
	consumer := NewConsumer(1, cond, messageChannel)

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := range 10 {
			producer.Produce(fmt.Sprintf("Message %d", i))
		}
	}()

	go func() {
		defer wg.Done()

		for range 10 {
			consumer.Consume()
		}
	}()

	wg.Wait()
}
