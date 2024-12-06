/*
读写锁：适合读多于写的并发场景，Lock()和Unlock()来操作写锁，RLock()和RUnlock()来操作读锁
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x     int64
	wg    sync.WaitGroup
	mtx   sync.Mutex
	rwMtx sync.RWMutex
)

func writeWithMtx() {
	mtx.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 写操作时间
	mtx.Unlock()
	wg.Done()
}

func readWithMtx() {
	mtx.Lock()
	time.Sleep(1 * time.Millisecond) // 读操作时间
	mtx.Unlock()
	wg.Done()
}

func writeWithRWMtx() {
	rwMtx.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwMtx.Unlock()
	wg.Done()
}

func readWithRWMtx() {
	rwMtx.RLock()
	time.Sleep(1 * time.Millisecond)
	rwMtx.RUnlock()
	wg.Done()
}

// 进行并发读写操作
func doSomething(wf, rf func(), wn, rn int) {
	start := time.Now()
	// wn个写操作
	for i := 0; i < wn; i++ {
		wg.Add(1)
		go wf()
	}
	// rn个读操作
	for i := 0; i < rn; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%d cost_time:%v\n", x, cost)
}

func main() {
	doSomething(writeWithMtx, readWithMtx, 5, 100)     // 使用互斥锁：x:5 cost_time:1.6302718s
	doSomething(writeWithRWMtx, readWithRWMtx, 5, 100) // 使用读写锁：x:10 cost_time:109.0396ms
}
