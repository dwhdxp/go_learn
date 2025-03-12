package main

import (
	"fmt"
	"sort"
)

// 切片，类似于"动态数组"
func main() {
	// 1.定义：var v_name []type
	var a []int  // 未初始化，默认值为nil
	b := []int{} // 声明并初始化为空，值不为nil
	c := []int{1, 2, 3, 4, 5}
	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	fmt.Printf("a_type=%v b_type=%v c_type=%v\n", a, b, c)
	// 和nil比较判断一个切片是否初始化，但不能判断是否为空
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	// fmt.Println(c == b)  error 切片是引用类型，不能直接比较，只能跟nil比较

	// 长度、容量：len() cap()
	fmt.Printf("len_a=%d, cap_a=%d\n", len(a), cap(a))
	fmt.Printf("len_b=%d, cap_b=%d\n", len(b), cap(b))
	fmt.Printf("len_c=%d, cap_c=%d\n", len(c), cap(c))
	// 使用len(a) == 0 来判断切片是否为空
	if len(a) == 0 || len(b) == 0 {
		fmt.Println("a and b is empty")
	}

	// 2.初始化
	// 基于数组初始化：s := arr[start:end]
	fmt.Println("Init......................")
	a1 := [5]int{1, 2, 3, 4, 5}
	b1 := a1[:4] // 切片是对数组的引用
	fmt.Printf("b1=%v b1_type=%T len=%d capacity=%d\n", b1, b1, len(b1), cap(b1))

	// 基于切片初始化
	c1 := b1[:len(b1)-1]
	fmt.Printf("c1=%v c1_type=%T len=%d capacity=%d\n", c1, c1, len(c1), cap(c1))

	// make()初始化切片：make([]T, length, capacity)
	d1 := make([]int, 5, 10) // var d1 []int = make([]int, 5, 10)
	fmt.Printf("d1=%v d1_type=%T len=%d capacity=%d\n", d1, d1, len(d1), cap(d1))

	// 3.拷贝
	// 赋值拷贝——浅拷贝，两者共享同一个底层数组
	fmt.Println("copy................")
	b2 := b1
	fmt.Printf("b2=%v\n", b2)
	b2[0] = 100
	b1[1] = 200
	fmt.Printf("b1=%v b2=%v\n", b1, b2) // 直接赋值，相当于浅拷贝，b2和b1共享同一个底层数组
	b1[0] = 1
	b1[1] = 2

	// copy()拷贝——深拷贝
	c2 := make([]int, 4, 10)
	copy(c2, b1)
	fmt.Printf("c2=%v\n", c2)
	c2[0] = 1000
	fmt.Printf("b1=%v c2=%v\n", b1, c2) // 修改c2不会影响b1

	// 4. 向切片动态添加元素，添加元素时可能涉及到动态扩容，因此一般使用原变量接收append()返回值
	fmt.Println("append................")
	a3 := []int{1, 2}
	fmt.Printf("a3=%v len=%d capacity=%d\n", a3, len(a3), cap(a3))
	a3 = append(a3, 3, 4, 5) // 向切片追加多个元素
	fmt.Printf("a3=%v len=%d capacity=%d\n", a3, len(a3), cap(a3))
	a3 = append(a3, a3...) // 向切片追加切片
	fmt.Printf("a3=%v len=%d capacity=%d\n", a3, len(a3), cap(a3))

	/*
		添加元素涉及到动态扩容
		扩容策略：old < 1024, new = 2*old；old > 1024, new = 1.25*old
	*/
	var c3 []int
	for i := 0; i < 10; i++ {
		c3 = append(c3, i)
		fmt.Printf("%v %d %d ptr=%p\n", c3, len(c3), cap(c3), c3) // 切片是引用类型，本质上是对数组的封装
	}

	// 在a的index之后插入元素：append(a[:index], append(b, a[index:]...)...)
	b3 := []int{5, 6, 7}
	b3 = append(b3[:1], append([]int{100, 100}, b3[1:]...)...)
	fmt.Printf("b3=%v\n", b3)

	// 5.删除切片中的元素：append(a[:index], a[index+1:]...)：删除下标为index的元素，相当于在index之前（跳过index）添加index之后的元素
	a4 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	a4 = append(a4[:2], a4[3:]...) // [0, 1, 3, 4, 5, 6, 7]
	fmt.Printf("%v\n", a4)
	a4 = a4[:len(a4)-2] // [0, 1, 3, 4, 5]
	fmt.Println(a4)

	// Test 1
	s := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)

	// Test 2
	var num = [...]int{9, 8, 4, 3, 6, 5, 1}
	sort.Ints(num[:])
	fmt.Println(num)
}
