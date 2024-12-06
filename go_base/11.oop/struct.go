package main

import "fmt"

// 类型别名 type newname typename
type myInt int

/*
	定义结构体：
	1.类名、属性名和方法名大小写：首字母大写表示公有，其他包可以访问；首字母小写，表示私有，只能在本包内访问
	2.定义方法时，如果不是this *T，那么使用时传入的是对象的拷贝，在函数内部的任何修改不会影响实际对象
*/
type User struct {
	Name  string
	Age   int
	level int
}

// 定义结构体的方法：相当于类内定义的成员函数
func (this User) Show() {
	fmt.Printf("User=%v\n", this)
}

func (this *User) SetName(newname string) {
	this.Name = newname
}

func (this *User) GetName() string {
	return this.Name
}

func main() {
	var a myInt = 10
	fmt.Println(a)

	user1 := User{Name: "lsx", Age: 25, level: 100}
	fmt.Printf("Name=%s\n", user1.GetName())
	user1.SetName("lxz")
	user1.Show()
}
