package main

import "fmt"

type Human struct {
	Name string
	Sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.walk()...")
}

// 继承：在子类中加上父类类名即可
type Superman struct {
	Human
	Level int
}

// 重写父类方法
func (this *Superman) Eat() {
	fmt.Println("Superman.eat()...")
}

func (this *Superman) Fly() {
	fmt.Println("Superman.fly()...")
}

func main() {
	h := Human{Name: "lsx", Sex: "female"}
	h.Eat()
	h.Walk()

	// var s Superman
	// s.Name = "lxz"
	// s.Sex = "female"
	// s.Level = 100
	s := Superman{Human{Name: "lxz", Sex: "female"}, 100}
	s.Eat()  // 调用子类方法
	s.Walk() // 调用父类方法
	s.Fly()
}
