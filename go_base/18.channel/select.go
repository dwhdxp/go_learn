/*
1.select可以同时监控和处理多个channel的收发操作；
2.若多个case同时满足，会随机选择一个执行；
3.没有case的select会一直阻塞
*/
package main

import "fmt"

func main() {
	c := make(chan int, 1)

	// 通过select实现对于channel的读写
	for i := 0; i < 10; i++ {
		select {
		// c可读
		case data := <-c:
			fmt.Println(data)
		// c可写
		case c <- i:
			fmt.Printf("send %d to ch\n", i)
		}
	}
}
