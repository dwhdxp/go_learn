package main

import "fmt"

// 定义单返回值函数
func foo1(a string, b int) int {
	fmt.Println("---- foo1 -----")
	fmt.Println(a)
	fmt.Println(b)

	c := 100
	return c
}

// 定义多返回值函数
// 方法一：(T, T...)
func foo2(a string, b int) (int, int) {
	fmt.Println("---- foo2 -----")
	fmt.Println(a)
	fmt.Println(b)

	r1 := 6
	r2 := 7
	return r1, r2
	// return 6, 7
}

// 方法二：(t1 T, t2 T...)
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- foo3 -----")
	fmt.Println(a)
	fmt.Println(b)

	// r1 r2为函数foo3的形参，初始默认值为0
	r1 = 6
	r2 = 7

	return
}

// 方法三：若类型相同，(t1, t2... T)
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- foo4 -----")
	fmt.Println(a)
	fmt.Println(b)

	return
}

func main() {
	c1 := foo1("aaa", 1)
	fmt.Printf("c1 = %d\n", c1)

	r1, r2 := foo2("bbb", 2)
	fmt.Printf("r1 = %d, r2 = %v\n", r1, r2)

	r1, r2 = foo3("ccc", 3)
	fmt.Printf("r1 = %d, r2 = %v\n", r1, r2)

	r1, r2 = foo4("ccc", 3)
	fmt.Printf("r1 = %d, r2 = %v\n", r1, r2)
}
