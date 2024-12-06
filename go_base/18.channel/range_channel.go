package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("sub goroutine over...")
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("sub goroutine write i=%d\n", i)
		}
		close(c)
	}()

	// 使用range循环从channel中接收数据
	for data := range c {
		fmt.Printf("main goroutine read data=%d\n", data)
	}
	defer fmt.Println("main goroutine over...")

}
