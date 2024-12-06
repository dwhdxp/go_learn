/*
1.channel可能在多个任务函数中被使用，通常我们会选择在不同的任务函数中对通道的使用进行限制，
比如限制通道在某个函数中只能执行发送或只能执行接收操作，这时就需要单向channel;
2.双向channel可以转换为单向channel，但是反过来不可以
*/
package main

import "fmt"

// 生产者：返回接收channel，限制其他模块向channel写数据
func Producer() <-chan int {
	c := make(chan int, 3)
	go func() {
		defer close(c)
		defer fmt.Println("Producer sub goroutine over")
		for i := 0; i < 10; i++ {
			if i&1 == 1 {
				c <- i
			}
		}
	}()
	return c
}

// 消费者：参数为接收channel
func Consumer(c <-chan int) {
	defer fmt.Println("Consumer over")
	for data := range c {
		fmt.Printf("Consumer print data:%d\n", data)
	}
}

func main() {
	// 只能用于写的管道：var send_c chan<- int
	// 只能用于读的管道：var read_c <-chan int
	c := Producer()
	Consumer(c)
}
