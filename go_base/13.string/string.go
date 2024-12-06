package main

import (
	"fmt"
	"strings"
)

// string：不可变只读的字符数组
// string常用操作
func main() {
	// 1.长度：len()
	s1 := "dwh"
	fmt.Println(len(s1))

	// 2.拼接：+ or fmt.Sprintf()
	s2 := s1 + "dxp"
	fmt.Println(s2)
	s3 := fmt.Sprintf("%s and %s", s1, s2)
	fmt.Println(s3)

	// 3.分割：strings.Split
	fmt.Println(strings.Split(s3, " "))
	fmt.Println(strings.Split(s3, "and"))
	fmt.Println(s3)
	fmt.Printf("%T \n", strings.Split(s3, " ")) // []string

	// 4.判断是否包含：strings.Contains
	fmt.Println(strings.Contains(s3, "and"))

	// 5.判断前后缀 strings.HasPrefix;strings.HasSuffix
	fmt.Println(strings.HasPrefix(s3, "dwh"))
	fmt.Println(strings.HasSuffix(s3, "dwh"))

	// 6.判断子串位置 strings.Index();strings.LastIndex()
	fmt.Println(strings.Index(s3, "lsx"))
	fmt.Println(strings.LastIndex(s3, "lsx"))

	// 7.join操作 strings.Join(a[]string, sep string) 将字符串切片a用sep字符串连接起来
	s4 := []string{"dwh", "lsx", "tb", "dxp"}
	fmt.Println(s4)
	fmt.Println(strings.Join(s4, " and "))

	// 8.修改字符串：不能直接修改string[index]，需要将其先转换为[]rune or []byte，完成后再转换为string
	s5 := "ttt"
	// s5[2] = 'b' error
	byteS5 := []byte(s5)
	byteS5[2] = 'b'
	s5 = string(byteS5)
	fmt.Println(s5)

	// 9. 拷贝
	s6 := "dwh dxp lsx"
	var s7 string = s6
	fmt.Println(s6, s7)
}
