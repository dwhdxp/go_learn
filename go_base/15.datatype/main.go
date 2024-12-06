package main

import (
	"fmt"
	"math"
)

/*
	1.Println() 打印字符串和变量，自动换行
		Println(a) √
		Println(str) √
	2.Printf()  格式化输出的字符串和变量，不自动换行
		fmt.Printf("hello %s", "world")
	3.Sprintf() 格式化输出的字符串和变量，返回字符串
		s := fmt.Sprintf("hello %s", "world")
*/

// 数据类型
func main() {
	// 1.整型 uint、int 8、16、32、64
	var num uint8 = 12
	// var num uint8 = 256 overflows
	fmt.Printf("%d \n", num)

	// 十进制 %d
	var a1 int32 = 5
	fmt.Printf("%b \n", a1)

	// 八进制 以0开头 %o
	var b1 int32 = 076
	fmt.Printf("%o \n", b1)
	fmt.Printf("%d \n", b1)

	// 十六进制 以0x开头 %x %X
	var c1 int32 = 0xfe
	fmt.Printf("%x \n", c1) // 小写
	fmt.Printf("%X \n", c1) // 大写
	fmt.Printf("%d \n", c1)

	// 2.浮点型 float32 float64
	var a2 float32 = 3.14
	fmt.Printf("%f \n", a2)
	fmt.Printf("%f \n", math.MaxFloat32)
	fmt.Printf("%f \n", math.MaxFloat64)

	// 3.布尔值：Go不允许将整型转换为布尔型；布尔值无法参与数值运算
	var a3 bool
	fmt.Println(a3)
	a3 = true
	fmt.Println(a3)

	// 4.字符串类型，打印一个网址 c:\code\go_learn\go.exe
	// fmt.Println("str := c:\code\go_learn\go.exe") error 存在转义字符'\'
	fmt.Println("str := c:\\code\\go_learn\\go.exe")
	// 用``定义多行字符串，不允许转义进行原样输出
	str1 := ` str 
		c:\code\go_learn\go.exe
		/
		//
	`
	fmt.Println(str1)

	/*
		5.字符型
		byte：==uint8, 代表ASCII码字符，处理ASCII字符串
		rune：==int32, 代表Unicode字符，处理中英文复合文本
	*/
	var a5 rune = '中'
	var b5 byte = 'x'
	fmt.Printf("%T %T\n", a5, b5)

	str2 := "dxp"
	for i := 0; i < len(str2); i++ { // s[i] is byte
		fmt.Printf("%c ", str2[i])
	}
	fmt.Println()
	str3 := "tb糖宝"
	for _, c := range str3 { // c is rune
		fmt.Printf("%c ", c)
	}
	fmt.Println()

	// 6.类型转换：只存在强制类型转换T()，不存在隐式转换
	a6 := 3.14
	fmt.Printf("%.3f \n", a6)
	fmt.Println(a6)
	fmt.Println(int(a6))
}
