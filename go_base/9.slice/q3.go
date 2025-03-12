package main

import "fmt"

func changeSlice(s2 []int) {
	s2[0] = -1
}

func changeSlice2(s3 []int) {
	// 动态扩容后，局部方法中的副本地址改变
	s3 = append(s3, 10, 11, 12, 13)
	fmt.Printf("address of s3 func:%p\n", s3)
	s3[0] = 100
}

func main() {
	// 越界问题：只要index≥len就会越界
	// 1.test7
	s := make([]int, 10, 12)
	//v := s[10] × 越界
	//fmt.Printf("%v\n", v)

	// 2.test8
	s1 := append(s[8:], []int{10, 11, 12}...)
	//v1 := s[10] × 越界：s1扩容后迁移到新地址，此时访问s[10]依然越界
	//fmt.Printf("%v\n", v1)
	fmt.Printf("test8 : len:%d cap:%d %v\n", len(s1), cap(s1), s1)
	fmt.Printf("test8 : len:%d cap:%d %v\n", len(s), cap(s), s)

	s11 := append(s[8:], 10, 11)
	//v1 := s[10] × 越界：s1未扩容，物理上s[10]等于10，但是s和s1仅array共享，len是独立的，所以10≥len(s)仍然越界
	//fmt.Printf("%v\n", v1)
	fmt.Printf("test8 : len:%d cap:%d %v\n", len(s11), cap(s11), s11)
	fmt.Printf("test8 : len:%d cap:%d %v\n", len(s), cap(s), s)

	// 传递与拷贝
	// 3.test9
	s2 := s[8:]
	changeSlice(s2) // 会影响s，共享内存空间
	fmt.Printf("test9 : len:%d cap:%d %v\n", len(s2), cap(s2), s2)
	fmt.Printf("test9 : len:%d cap:%d %v\n", len(s), cap(s), s)

	// 4.test10
	s[8] = 0
	s3 := s[8:]
	changeSlice2(s3)
	fmt.Printf("test10 : len:%d cap:%d address:%p %v\n", len(s3), cap(s3), s3, s3)
	fmt.Printf("test10 : len:%d cap:%d address:%p %v\n", len(s), cap(s), s, s)

	// 5.test11
	s4 := []int{0, 1, 2, 3, 4}
	s4 = append(s4[:2], s4[3:]...)
	fmt.Printf("test11 : len:%d cap:%d %v\n", len(s4), cap(s4), s4)
	//v := s4[4]  × 越界
	//fmt.Println(v)

	// 追加
	// 6.test12
	s5 := make([]int, 512)
	s5 = append(s5, 1)
	fmt.Printf("test12 : len:%d cap:%d %v\n", len(s5), cap(s5), s5)

	// 7.test13
	// s := make([]int, 10, 12)
	s6 := s[3:]
	s6 = append(s6, 10)
	fmt.Printf("test13 : len:%d cap:%d %v\n", len(s6), cap(s6), s6)
	fmt.Printf("test13 : len:%d cap:%d %v\n", len(s), cap(s), s)
}
