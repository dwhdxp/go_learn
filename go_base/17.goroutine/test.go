package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		// 由于闭包和并发：导致结果可能有很多个5
		go func() {
			fmt.Println(i) // i是外层作用域变量，可能sub还未打印时，i已经加到x了
		}()
	}

	/*
		想要并发打印0-4，可以通过给函数传参，避免闭包行为
		for i := 0; i < 5; i++ {
			go func(i int) {
				fmt.Println(i)
			}(i)
		}
	*/

	time.Sleep(time.Second)
}
