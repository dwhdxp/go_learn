/*
1.sync.WaitGroup可以实现阻塞等待一组goroutine并发操作完成，其内部维护了一个计数器，来控制并发任务的同步；
2.Add(i):让计数器 +i；Done():让计数器 -1；Wait():阻塞等待直到计数器为0时，解除阻塞，表示所有并发任务完成；
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine over 就让计数器-1
	fmt.Println("hello", i)
}

func main() {
	defer fmt.Println("main goroutinue over")
	for i := 0; i < 5; i++ {
		wg.Add(1) // 启动一个新的goroutine就让计数器 + 1
		// 每次执行打印顺序不一致，主要由于go的调度是随机的
		go hello(i)
	}
	wg.Wait()
}
