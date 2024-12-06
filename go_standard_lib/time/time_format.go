// 时间格式化必须用：2006(y) 01(m) 02(d) 15(h) 04(m) 05(s)
package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间格式化
	now := time.Now()
	fmt.Println(now)

	t1 := now.Format("2006-01-02")
	fmt.Println(t1)
	t2 := now.Format("2006.01.02 15:04:05")
	fmt.Println(t2)
	t3 := now.Format("2006-01-02 15:04:05")
	fmt.Println(t3)

	// 根据时区去解析一个字符串格式的时间
	// 1.获取时区
	timeStr := "2024/12/06 22:00:00"
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("get location time failed, err:%v\n", err)
		return
	}
	// 2.按时区解析字符串格式时间：ParseInLocation()
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("parse timeStr failed, err:%v\n", err)
		return
	}
	// 不按时区解析：Parse()
	timeObj1, err1 := time.Parse("2006/01/02 15:04:05", timeStr)
	if err1 != nil {
		fmt.Printf("parse timeStr failed, err:%v\n", err1)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj1)

}
