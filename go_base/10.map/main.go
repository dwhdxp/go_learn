package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1.初始化：字面量初始化、make初始化
	// map[KeyType]ValueType{}
	fmt.Println("init...........................")
	scoreMap := map[string]int{
		"dwh": 100,
		"lsx": 99,
	}
	scoreMap["tb"] = 1 // 访问的key不存在，会直接插入
	fmt.Println(scoreMap)
	fmt.Printf("the score of dwh=%d\n", scoreMap["dwh"])
	fmt.Println(scoreMap["dwh"])
	fmt.Printf("type=%T\n", scoreMap)

	// make(map[KeyType]ValueType, capacity)
	userInfo := make(map[string]string, 5)
	fmt.Println(userInfo)
	userInfo["dwh"] = "he"
	userInfo["lsx"] = "she"
	userInfo["tb"] = "it"
	fmt.Println(userInfo)

	/*
		map是引用类型，未初始化或为nil的map，可以访问但是无法存放元素，必须分配空间
		var mp map[string]int
		mp["ttt"] = 2
		fmt.Println(mp) // panic: assignment to entry in nil map
	*/

	// 2.判断key是否存在
	fmt.Println("exit...........................")
	value, exist := scoreMap["dxp"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("key is not exist")
		fmt.Println(scoreMap)
	}

	// 求长度：len()
	fmt.Println(len(scoreMap))
	fmt.Println(len(userInfo))

	// 3.遍历：map是无序的，遍历顺序与插入键值对顺序无关
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
	// 只遍历key
	for key := range scoreMap {
		fmt.Println(key)
	}

	// 4.删除：delete(map, key)；清空：clear(map)
	fmt.Println("delete、clear...........................")
	delete(scoreMap, "tb")
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
	clear(scoreMap)
	fmt.Println(scoreMap)

	fmt.Println("apply...........................")
	// 5.按指定顺序遍历map
	scoreMap["tb"] = 2
	scoreMap["txb"] = 2
	scoreMap["dwh"] = 100
	scoreMap["lsx"] = 99
	scoreMap["dxp"] = 10
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 10)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for i := 0; i < len(keys); i++ {
		fmt.Println(keys[i], scoreMap[keys[i]])
	}
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// 6.定义切片中的元素为map类型
	var mapSlice = make([]map[string]int, 3)
	for index, value := range mapSlice {
		fmt.Printf("index=%d, value=%v\n", index, value)
	}
	// map必须开辟空间，才能存储键值对
	mapSlice[0] = make(map[string]int, 5)
	mapSlice[1] = make(map[string]int, 5)
	mapSlice[0]["dwh"] = 100
	mapSlice[0]["dxp"] = 99
	mapSlice[1]["lsx"] = 100
	mapSlice[1]["lxz"] = 99
	for index, value := range mapSlice {
		fmt.Printf("index=%d, value=%v\n", index, value)
	}

	// 7.map中值为切片类型
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key1 := "dwh"
	key2 := "lsx"
	val1, exist1 := sliceMap[key1]
	val2, exist2 := sliceMap[key2]
	if !exist1 {
		val1 = make([]string, 0, 2)
		val1 = append(val1, "he", "bendan")
		sliceMap[key1] = val1
	}
	if !exist2 {
		val2 = make([]string, 0, 2)
		val2 = append(val2, "she", "xiaozhu")
		sliceMap[key2] = val2
	}
	fmt.Println(sliceMap)
}
