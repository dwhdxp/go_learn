package main

import (
	"fmt"
	"time"
)

// 统计一段代码的执行耗时时间，单位精确到微秒
func CountTimeDur() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	sub := time.Now().Sub(start)
	fmt.Printf("sub time:%v\n", sub/1000)
}

func main() {
	// 获取当前时间，格式化输出为2017/06/19 20:30:05格式
	now := time.Now()
	t := now.Format("2006/01/02 15:04:05")
	fmt.Println(t)

	CountTimeDur()
}
