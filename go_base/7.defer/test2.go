package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	// defer函数的参数值会全部确定
	defer calc("AA", x, calc("A", x, y)) // defer calc("AA", 1, 3)
	x = 10
	defer calc("BB", x, calc("B", x, y)) // defer calc("BB", 10, 12)
	y = 20
}
