/*
1.sync.Map：并发安全的map（原生的map不支持并发操作），接收的interface{}类型，无需指定键值对类型
2.Store()：存储键值对；Load()：查询key对应的value；Delete()：删除key；
*/
package main

import (
	"fmt"
	"sync"
)

var smp = sync.Map{}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			smp.Store(i, i+100)
			v, _ := smp.Load(i)
			fmt.Printf("k:%v v:%v\n", i, v)
		}(i)
	}
	wg.Wait()
}
