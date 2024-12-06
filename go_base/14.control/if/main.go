package main

import "fmt"

// if
func main() {
	/*
		if expression1 {

		}else if expression2 {

		}else {

		}
	*/
	a := 5
	if a > 5 {
		fmt.Println(a)
	} else if a < 0 {
		a++
	} else {
		a--
	}
	fmt.Println(a)

	if b := 3; b < 1 {
		b++
	} else {
		b--
	}
	// fmt.Println(b) error; b是在if中定义的，在if之外无法使用
}
