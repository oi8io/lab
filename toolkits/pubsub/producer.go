package pubsub

import (
	"context"
	"fmt"
)

// Producer 生产者：生成整数的倍数序列
func Producer(factor int, out chan<- int) {
	context.Background()
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// Consumer 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}


