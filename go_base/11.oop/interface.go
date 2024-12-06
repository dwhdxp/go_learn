package main

import "fmt"

/*
	多态：1.定义父类接口；2.子类重写父类所有接口；3.父类指针指向子类对象，调用接口
*/
type Animal interface {
	Sleep()
	GetColor() string
	GetAnimalType() string
}

// cat子类
type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	fmt.Println(this.Color)
	return this.Color
}

func (this *Cat) GetAnimalType() string {
	return "Cat"
}

// dog子类
type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	fmt.Println(this.Color)
	return this.Color
}

func (this *Dog) GetAnimalType() string {
	return "Dog"
}

func showAnimal(animal Animal) {
	fmt.Println("the type of animal is", animal.GetAnimalType())
}

/*
	interface{}类型：空接口，一种万能类型，可以接收任意类型的函数参数
	类型断言：1.value := arg.(type_name)，断言成功，value得到正确的值；断言失败，会导致panic的错误；
					 2.value, ok := arg.(T)，为了避免panic错误，可以使用ok来判断是否断言成功
*/
// 能接受任意类型
func myFunc(arg interface{}) {
	value, ok := arg.(string)
	if !ok {
		fmt.Println("panic:the type is not string")
	} else {
		fmt.Println(value)
	}
}

func main() {
	var animal Animal
	// 父类指针指向子类
	animal = &Cat{Color: "black"}
	animal.Sleep()
	animal.GetColor()
	animal.GetAnimalType()
	showAnimal(animal)

	animal = &Dog{Color: "green"}
	animal.Sleep()
	animal.GetColor()
	animal.GetAnimalType()
	showAnimal(animal)

	fmt.Println("interface.........")
	myFunc("lsx")
	myFunc(1)
	myFunc(1.1)
	myFunc(false)
	myFunc(animal)
}
