package main

import "fmt"

/*
	1.defer在当前函数生命周期全部结束后才执行，即return后执行defer；
	2.函数定义多个defer，它们会以FILO(栈)顺序执行
	3.defer函数所有的参数都需要提前确定其值，因为当defer执行时，函数生命周期已经结束
*/
func deferFunc1() int {
	fmt.Println("A")
	return 0
}

func deferFunc2() int {
	fmt.Println("B")
	return 0
}

func deferFunc3() int {
	fmt.Println("C")
	return 0
}

func returnFun() int {
	fmt.Println("return func call ...")
	return 0
}

func fun() int {
	defer deferFunc1()
	defer deferFunc2()
	defer deferFunc3()

	return returnFun()
}

func main() {
	fun()
}
