package main

import "fmt"

func main() {
	// 定义：必须初始化，格式将var替换成const
	const pi float64 = 3.14
	const e = 2.7
	const (
		E   = 2.623
		pai = 3.1415
	)
	fmt.Println(pi, e, pai, E)

	// const定义枚举类型
	// iota：常量计数器，只能在常量表达式中使用，每行的iota会累加1，第一行的iota默认值为0
	const (
		a = iota // iota = 0
		b        //   iota = 1
		c        //   iota = 2
	)
	fmt.Println(a, b, c)

	const (
		a1 = iota // iota = 0
		_         // iota = 1

		b1 = 100  // iota = 2
		c1 = iota // iota = 3
		d1        // iota = 4
	)
	const e1 = iota
	fmt.Println(a1, b1, c1, d1, e1)

	const (
		a2, b2 = iota + 1, iota + 2 // iota = 0
		c2, d2 = iota * 2, iota * 3 // iota = 1
		e2, f2                      // iota = 2
	)
	fmt.Println(a2, b2, c2, d2, e2, f2)

	// 常用于定义数量级
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println(KB, MB, GB, TB)
}
