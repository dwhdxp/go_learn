/*
atomic：针对整数类型提供了原子操作来保证并发安全
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x  int64
	wg sync.WaitGroup
)

func AtomicAdd() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go AtomicAdd()
	}
	wg.Wait()
	fmt.Println("x =", x)
}
