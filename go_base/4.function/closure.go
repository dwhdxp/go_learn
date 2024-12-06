package main

import "fmt"

// 返回值为函数类型的函数
func foo() func() {
	data := 10
	// 返回的函数调用时，找不到外层变量data，就会到外层函数(foo1)找data
	return func() {
		fmt.Printf("return func data:%d\n", data)
	}
}

func adder() func(int) int {
	x := 10
	return func(y int) int {
		x += y
		return x
	}
}

func adder1(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	// 闭包 = 函数 + 外层变量的引用
	r := foo() // f就是一个闭包，不仅包含函数，还包含该函数引用的外层变量
	r()

	// 在f的生命周期内，外层变量x一直有效
	f := adder()
	fmt.Println(f(10)) // 20
	fmt.Println(f(20)) // 40
	fmt.Println(f(30)) // 70

	f1 := adder1(10)
	fmt.Println(f1(10)) // 20
	fmt.Println(f1(20)) // 40
	fmt.Println(f1(30)) // 70

	f2 := adder1(20)
	fmt.Println(f2(10)) // 30
	fmt.Println(f2(20)) // 50
	fmt.Println(f2(30)) // 80

}
