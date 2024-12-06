/*
1.recover()必须搭配defer使用，panic会直接终止程序，终止后会执行defer
2.defer必须在出现panic的位置之前定义，否则会导致无法被注册，panic出现后程序会直接崩溃，程序不会注册defer
*/
package main

import "fmt"

func a() {
	fmt.Println("print AA")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b func")
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("func b error")
	// 	}
	// }()
}

func c() {
	fmt.Println("print CC")
}

func main() {
	a()
	b()
	c()
}
