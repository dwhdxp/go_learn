package main

import (
	"fmt"
	"reflect"
)

/*
	1.反射就是在运行时可以动态的获取一个接口变量的类型信息和值信息;
	2.reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type;
		TypeOf：Name()求type，Kind()求底层类型（一般用于区别指针、结构体）
		ValueOf：返回的是reflect.Value类型，可以与原始值相互转换，先通过Interface()转换为interface{}类型，
						 然后再通过类型断言转换为指定类型
	3.结构体反射：(1)通过reflect.Type的NumField()和Field(i)可以分别获取结构体中字段的数量和第i个字段
							 Field返回的是一个StructField类型，类型中定义了Name、Type、Tag等字段表示信息;
               (2)通过reflect.Type的NumMethod()和Method(i)可以获取结构体中方法的数量和第i个方法
							 通过反射调用方法时，参数必须是[]reflect.Value类型
*/

func justVar(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("var_type=%v var_kind=%v value=%v\n", t.Name(), t.Kind(), v)
	// 转换
	realVal, ok := v.Interface().(string)
	if !ok {
		fmt.Println("type is not string")
	} else {
		fmt.Println("type is string")
		fmt.Println(realVal)
	}

}

type User struct {
	Name string
	Id   int
	Age  int
}

func (user User) Show(name string, id int, age int) {
	fmt.Println("show user method")
}

func (user User) Call(name string, id int, age int) {
	fmt.Printf("call user method:%s %d %d\n", name, id, age)
}

func reflectStruct(input interface{}) {
	// 获取input的type和val
	inputType := reflect.TypeOf(input)
	inputVal := reflect.ValueOf(input)
	fmt.Printf("inputType=%v inputKind=%v inputVal=%v\n", inputType.Name(),
		inputType.Kind(), inputVal)

	// 获取结构体内字段：NumField()获取字段数量;Field(i)获取其字段;最后通过Field的Interface()得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		val := inputVal.Field(i).Interface()
		fmt.Printf("%s  %v:%v\n", field.Name, field.Type, val)
	}

	fmt.Println("Struct method................")
	// 获取结构体内方法并调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s  %v\n", method.Name, method.Type)
		// 通过反射调用方法：[]reflect.Value{}，可以无参、有参
		args := []reflect.Value{reflect.ValueOf("lsx"), reflect.ValueOf(2), reflect.ValueOf(99)}
		inputVal.Method(i).Call(args)
	}
}

func main() {
	s := "dxp"
	justVar(s)
	justVar(&s)

	fmt.Println("Struct reflect..................")
	user := User{Name: "dxp", Id: 1, Age: 18}
	reflectStruct(user)
}
