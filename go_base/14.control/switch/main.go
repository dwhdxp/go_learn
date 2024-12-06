package main

import "fmt"

func main() {
	/*
		switch expr {
		case case1:
			statement1
		case case2:
			statement2
		default:
			default statement
		}
	*/
	score := 85
	switch {
	case score > 0 && score < 60:
		fmt.Println("C")
	case score >= 60 && score < 70:
		fmt.Println("B")
	case score >= 70 && score < 80:
		fmt.Println("A")
	case score >= 80 && score <= 100:
		fmt.Println("S")
	default:
		fmt.Println("error")
	}
}
