package main

import "fmt"

// 数组
func main() {
	// 1.定义并初始化：var v_name [N]type，定长且长度只能是常量表达式
	var nums1 [5]int
	nums2 := [6]int{1, 2, 3, 4, 5, 6}
	nums3 := new([7]int)          // 获得一个指向数组的指针
	nums4 := [...]int{1, 2, 3, 4} // 自动推断数组长度
	fmt.Println(nums1, nums1[1])
	fmt.Println(nums2, nums2[1])
	fmt.Println(nums3, nums3[1])
	fmt.Println(nums4, nums4[1])

	// 2.内置函数：cap() len()
	fmt.Printf("length=%d nums1_type=%T\nlength=%d nums2_type=%T\n", len(nums1), nums1, len(nums2), nums2)
	fmt.Printf("length=%d nums3_type=%T\nlength=%d nums4_type=%T\n", len(nums3), nums3, len(nums4), nums4)
	fmt.Printf("capacity=%d\n", cap(nums1)) // 数组的容量和长度相同

	// 3.切割：切割数组的格式为arr[startIndex:endIndex]，切割的区间为[startIndex, endIndex)
	fmt.Println(nums2[1:])  // [1, 6)
	fmt.Println(nums2[:6])  // [0, 6)
	fmt.Println(nums2[1:3]) // [1, 3)
	fmt.Println(nums2[2:3]) // [2, 3)

	// 4.遍历：for or for_range
	for i := 0; i < len(nums2); i++ {
		fmt.Printf("%d ", nums2[i])
	}
	fmt.Println()
	for index, value := range nums2 {
		fmt.Println(index, value)
	}

	// 5.二维数组
	nums5 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for _, x := range nums5 {
		for _, y := range x {
			fmt.Printf("%d ", y)
		}
		fmt.Println()
	}

	// 6.Test：找出数组中和为target的两个元素的下标
	target := 8
	for i := 0; i < len(nums2); i++ {
		for j := i + 1; j < len(nums2); j++ {
			if nums2[i]+nums2[j] == target {
				fmt.Println(i, j)
			}
		}
	}

}
