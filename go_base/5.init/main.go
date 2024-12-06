package main

/*
	1.go中有两个保留的函数init()和main(),go程序会自动调用init、main函数
	2.执行顺序：编译main包 -> 依次进入main包中import的其他包 -> 对包中包级的常量和变量进行初始化 -> other.init()
	         -> 返回main包中 -> 初始化main包中包级的常量和变量 -> main.init() -> main.main()
	3.如果一个包会被多个包同时导入，那么它只会被导入一次
*/

import (
	"fmt"

	"project/lib1"
	"project/lib2"
)

func init() {
	fmt.Println("libmain init ...")
}

func main() {
	lib1.Lib1Func()
	lib2.Lib2Func()
}
