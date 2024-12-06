/*
1.goroutine是轻量级用户线程，由go进行调度，并且go会动态的为其分配合适的栈空间(几KB)；
2.当需要让任务并发执行时，只需要将任务包装成一个函数，开启goroutine执行函数；
3.当main goroutine over时，所有main创建的sub goroutine都会over；
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func NewRoutine() {
	defer fmt.Println("new routine exit")
	i := 0
	for i < 5 {
		fmt.Printf("new routine print i=%d\n", i)
		time.Sleep(1 * time.Second)
		i++
		if i == 2 {
			// 直接终止当前goruntine
			runtime.Goexit()
		}

	}
}

func main() {
	// 有名函数创建goroutine
	go NewRoutine()

	// 匿名函数创建goroutine
	go func(a, b int) bool {
		defer fmt.Println("new2 routine eixt")
		fmt.Printf("new2 routine print a=%d b=%d\n", a, b)
		return true
	}(10, 10)

	time.Sleep(5 * time.Second)
	fmt.Println("main routine exit")
}
