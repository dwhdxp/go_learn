package main

import "fmt"

/*
	1.一个变量包括两个部分(type, value)，type包括static_type(int,string,struct...)和concrete_type(实际使用的类型);
	2.指定类型的变量的类型是静态的(static_type),在创建变量的时候就已经确定；而interface{}类型变量可以存储任意类型的变量
		在运行时动态调整type(concrete_type)，因此反射主要与interface{}变量相关;
	3.类型断言能否成功，取决于变量的concrete_type，而不是static_type，因此，一个reader变量如果它的concrete_type
		也实现了write方法的话，它也可以被类型断言为writer;
*/

type Reader interface {
	Read()
}

type Writer interface {
	Writer()
}

// 具体类型
type Book struct{}

func (this *Book) Read() {
	fmt.Println("read a book")
}

func (this *Book) Writer() {
	fmt.Println("write a book")
}

func main() {
	// b = pair<type:*Book,value:Book{}>
	b := &Book{}
	fmt.Printf("type = %T\n", b)

	// 创建接口变量时，r = pair<type:nil, value:nil>，因此可以通过 r == nil 来判断是否接口为空
	var r Reader
	// r = pair<type:*Book,value:Book{}>
	r = b
	fmt.Printf("type = %T\n", r)
	r.Read()

	// w = pair<type:nil, value:nil>
	var w Writer
	// w = pair<type:Book,value:Book{}>
	// reader的concrete_type是*Book，writer的也是，因此两者可以互相断言
	w = r.(Writer)
	fmt.Printf("type = %T\n", w)
	w.Writer()
}
