package main

import "fmt"

/*
	注意事项：
	1.声明的变量必须使用
	2.函数外的每个语句都必须以关键字var、const、func等开始
	3._多用于占位，表示忽略值
*/

func main() {
	// 变量声明格式: var 变量名 变量类型；
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)

	// 初始化：1.默认 int(0) string("") bool(false) func、指针、切片(nil) 2.var name string = "dwh"
	var name1 string = "dwh"
	var age1 int = 24
	var name2, age2 = "lsx", 25 // 多变量声明
	fmt.Println(name1, age1, name2, age2)

	// 类型推导：根据变量初始值自动类型推导
	var name3 = "xtb"
	var age3 = 2
	fmt.Println(name3, age3)
	fmt.Printf("type of name3 = %T, age3 = %T\n", name3, age3)

	// 短变量声明：省去var关键字，但只能在函数内使用，不能声明全局变量
	name4 := "Dwh"
	age4 := 24
	fmt.Println(name4, age4)

	// 批量声明：一般用来声明全局变量
	var (
		a string
		b int
		c float64
	)
	fmt.Println(a, b, c)

	var a1, b1, c1 = "dwh", 24, 3.14
	fmt.Println(a1, b1, c1)

	// 匿名变量：_表示，不占用命名空间，也不分配内存
	x, _ := 10, "lsx"
	_, y := 10, "lsx"
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
