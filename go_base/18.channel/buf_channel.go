package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		创建一个有缓冲的channel：make(chan type, buf_size)
		有缓冲channel：当channel满时，写goroutine会阻塞；当channel为空时，读goroutine会阻塞
	*/
	ch := make(chan int, 3)
	fmt.Printf("channel_len=%d channel_cap=%d\n", len(ch), cap(ch))

	go func() {
		defer fmt.Println("sub goroutine over")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Printf("i=%d channel_len=%d channel_cap=%d\n", i, len(ch), cap(ch))
		}
	}()

	defer fmt.Println("main goroutine over")
	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Printf("main print num=%d\n", num)
	}
	// 确保sub可以over
	time.Sleep(1 * time.Second)
}
