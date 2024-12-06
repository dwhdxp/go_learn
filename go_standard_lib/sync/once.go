/*
1.sync.Once：确保某操作在高并发场景下只执行一次，例如加载配置文件、进行初始化操作、闭关通道、并发安全单例模式等；
2.Once的Do()接收的是个无参的函数，因此如果想要使用参数需要通过闭包的方式；
3.sync.Once内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成；
*/
package main

import (
	"sync"
)

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	// 确保单例初始化操作只执行一次
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

func main() {

}
