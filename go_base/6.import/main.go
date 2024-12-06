package main

import (
	"fmt"
	_ "project/lib1" // _ 表示匿名导入，即实际没有使用导入包的内容
	_ "project/lib2"
)

func init() {
	fmt.Println("libmain init ...")
}

func main() {
	fmt.Println("libmain main()...")
}
