package main // 定义程序的包名，表明该文件属于哪个包;package main表示一个可独立执行的程序

import (
	"fmt"  // 导入fmt包，fmt包实现了格式化输入输出的功能
	"time" // 导入time包，time包提供了时间日期相关的功能
)

func main() {
	fmt.Println("Hello World from Go!")
	fmt.Println("The time is", time.Now()) // 输出当前时间

	time.Sleep(1 * time.Second) // 休眠1秒
}
