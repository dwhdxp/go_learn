/*
1.对一个关闭的通道再发送值就会导致 panic;
2.对一个关闭的通道进行接收会一直获取值直到通道为空;
3.关闭一个已关闭的channel会导致 panic
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("sub goroutine over")
		for i := 0; i < 4; i++ {
			c <- i
			// close(c)     向关闭的通道写数据会导致panic
		}
		// 关闭通道
		close(c)
	}()

	for {
		// 通道关闭时，ok会返回false，value为对应类型的零值；否则，ok会返回true
		if num, ok := <-c; ok {
			fmt.Printf("num=%d\n", num)
		} else {
			break
		}
	}

	fmt.Println("main goroutine over")
}
