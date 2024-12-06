package main

import "fmt"

/*
	注意：Go中无自增、自减运算符，它们被降级为语句statement，并且只能位于操作数的后侧
	a++
	a--
*/

// 运算符
func main() {
	// 1.算术运算符
	fmt.Println(1 + 2*3)
	fmt.Println(3 / 2)
	fmt.Println(3 % 2)

	// 2.关系运算符
	fmt.Println(5 > 4)
	fmt.Println(4 >= 5)

	// 3.逻辑运算符
	fmt.Println(5 > 4 && 4 >= 5)
	fmt.Println(5 > 4 || 4 >= 5)
	fmt.Println(!(5 > 4))

	// 4.位运算符 & | ^ << >> 对整数在内存中的二进制位进行操作
	a := 5                             // 101
	b := 1                             // 001
	fmt.Printf("%b %d \n", a&b, a&b)   // 001 1
	fmt.Printf("%b %d \n", a|b, a|b)   // 101 5
	fmt.Printf("%b %d \n", a^b, a^b)   // 100 4
	fmt.Printf("%b %d \n", a<<1, a<<1) // 1010 10
	fmt.Printf("%b %d \n", a>>1, a>>1) // 10 2

	// 5.赋值运算符
	c := 1
	c += a
	c *= b
	fmt.Println(c)
}
