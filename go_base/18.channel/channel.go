/*
1.channel是一种特殊类型，类型消息队列FIFO，用于多个goroutine 之间通信，同时也起到一定同步效果；
2.channel未初始化默认值为nil
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		创建一个无缓存channel：make(chan type)
		无缓冲channel(同步通道)：要求发送goroutine和接收goroutine同时准备好，才能完成发送和接收操作，
													  否则channel会阻塞先执行操作的goroutine，直至另一个goroutine也执行操作。
	*/
	ch := make(chan int)

	go func(a int) {
		defer fmt.Println("sub routine over")
		fmt.Printf("sub routine print a=%d\n", a)
		// time.Sleep(1 * time.Second)
		ch <- a
	}(100)

	defer fmt.Println("main routine over")
	// 无缓冲channel同步：main阻塞，无法进入channel，此时sub进入channel后会被锁住直至main进入channel
	time.Sleep(2 * time.Second)
	b := <-ch
	fmt.Printf("main routine print b=%d\n", b)
}
