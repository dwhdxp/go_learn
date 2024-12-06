/*
1.time.Time类型表示时间对象类型
2.时间戳和time对象可以相互转换：Unix()
3.时间间隔time.Duration:两个时间相减得到时间间隔，以nm为单位
*/
package main

import (
	"fmt"
	"time"
)

func GetTime() {
	// 1.时间对象time.Time
	fmt.Println("time.Time.........................")
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)

	// 2.时间戳：1970.1.1 到当前的总秒数
	fmt.Println("时间戳.........................")
	timeStamp1 := now.Unix()      // s 1733492257
	timeStamp2 := now.UnixMilli() // ms
	fmt.Println("time.Time-->时间戳:time.Time.Unix():", timeStamp1, timeStamp2)
	// 将时间戳转换为具体时间格式
	t := time.Unix(timeStamp1+3600, 0)
	fmt.Println("时间戳-->time.Time:time.Unix():", t)

	// 3.时间间隔 time.Duration
	fmt.Println("time.Duration.........................")
	n := 5
	time.Sleep(time.Duration(n) * time.Nanosecond) // 需要强制将int转换为time.Duration才可以使用
	time.Sleep(5 * time.Nanosecond)
	t1 := t.Add(2 * time.Hour)
	fmt.Println("time add 2h:", t1)
}

func main() {
	GetTime()
}
