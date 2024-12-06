package main

import "fmt"

func main() {
	/*
		for init statement; expression; post statement {
			execute statement
		}
	*/
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	/*
		Go语言中没有while，通过for忽略初始语句和后置语句来代替while
		for {
			execute statement
		}
	*/
	i := 10
	for i > 0 {
		fmt.Printf("%d ", i)
		i--
	}
	fmt.Println()

	// 打印9*9乘法表
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d*%d=%2d ", i, j, i*j)
		}
		fmt.Println()
	}
}
