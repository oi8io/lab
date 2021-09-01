package pubsub

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestNewPublisher(t *testing.T) {
	p := NewPublisher(5*time.Second, 5)
	defer p.Close()
	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("Hello World")
	p.Publish("Hello golang")
	p.Publish("Hello golang")
	p.Publish("Hello World")
	p.Publish("Hello golang")
	p.Publish("Hello golang")
	p.Publish("Hello World")
	p.Publish("Hello World")
	p.Publish("Hello golang")
	p.Publish("Hello World")
	p.Publish("Hello golang")
	p.Publish("Hello golang")

	go func() {
		for msg := range all {
			fmt.Println("All msg", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("Golang msg", msg)
		}
	}()

	time.Sleep(time.Second * 2)
}

func TestConsumer(t *testing.T) {
	ch := make(chan int, 64)
	go Producer(2,ch)
	go Producer(5,ch)
	go Consumer(ch)
	go Consumer(ch)
	time.Sleep(time.Second*1)
}