package main

import "fmt"

func main() {
	// 1.test4
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Printf("test4: len:%d cap:%d %v\n", len(s1), cap(s1), s1)

	// 2.test5
	s2 := s[8:9]
	fmt.Printf("test5: len:%d cap:%d %v\n", len(s2), cap(s2), s2)

	// 3.test6
	s3 := s[8:]
	s3[0] = -1
	fmt.Printf("test6: len:%d cap:%d %v\n", len(s3), cap(s3), s3)
	fmt.Printf("test6: len:%d cap:%d %v\n", len(s), cap(s), s)
}
