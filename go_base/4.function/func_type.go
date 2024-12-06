package main

import "fmt"

// 定义一个函数类型，类似可调用对象
type calculate func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func cal(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	var c calculate
	c = add
	fmt.Printf("type:%T res=%d\n", c, c(1, 9))
	// 匿名函数
	c = func(a, b int) int {
		return a * b
	}
	fmt.Printf("type:%T res=%d\n", c, c(2, 9))

	ans := cal(10, 10, sub)
	fmt.Println(ans)
}
