package main

import "fmt"

func main() {
	/**
	 * 老版本：oldcap<1024时，newcap=2*oldcap；oldcap>1024，newcap=1.25*oldcap
	 * go1.18扩容策略调整：oldcap<256时，newcap=2*oldcap；
	**/
	// 1.test1
	s1 := make([]int, 10)
	s1 = append(s1, 1)
	fmt.Printf("test1: len:%d cap:%d %v\n", len(s1), cap(s1), s1)

	// 2.test2
	s2 := make([]int, 0, 10)
	s2 = append(s2, 1)
	fmt.Printf("test2: len:%d, cap:%d %v\n", len(s2), cap(s2), s2)

	// 3.test3
	s3 := make([]int, 10, 11)
	s3 = append(s3, 1)
	fmt.Printf("test3: len:%d cap:%d %v\n", len(s3), cap(s3), s3)

	s4 := make([]int, 257)
	s5 := make([]int, 255, 258)
	s4 = append(s4, s5...)
	fmt.Printf("len:%d cap:%d\n", len(s4), cap(s4))
}
