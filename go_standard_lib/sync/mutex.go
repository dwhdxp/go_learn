package main

import (
	"fmt"
	"sync"
)

var (
	x   int64
	wg  sync.WaitGroup
	mtx sync.Mutex
)

func adder() {
	defer fmt.Println("sub goroutine over")
	for i := 0; i < 2000; i++ {
		mtx.Lock()
		x += 1
		mtx.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go adder()
	go adder()

	wg.Wait()
	fmt.Println("main goroutine over x = ", x)
}
