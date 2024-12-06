package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	// 捕获0
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	// 如果defer修改了命名返回值，会影响返回值；修改本地变量并不会影响返回的值；
	fmt.Println(f1()) // 5，修改本地参数
	fmt.Println(f2()) // 6，修改命名返回值，会影响返回值
	fmt.Println(f3()) // 5，修改本地参数
	fmt.Println(f4()) // 5，提前确定defer参数，修改内部参数不会修改明明返回值
}
