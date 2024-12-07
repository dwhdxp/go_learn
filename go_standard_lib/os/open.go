/*
1.文件类型为 *File
2.os.Open()：以只读的方式打开一个文件，可通过os.Close()关闭
3.os.OpenFile()：以指定模式和权限打开文件
*/
package main

import (
	"fmt"
	"io"
	"os"
)

// Open()
func OpenFile1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file %s failed, err:%v\n", filename, err)
		return
	}
	defer file.Close()
}

// OpenFile()
func OpenFile2(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file %s failed, err:%v\n", filename, err)
		return
	}
	defer file.Close()
}

// Read()
func ReadFile1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file %s failed, err:%v\n", filename, err)
	}
	defer file.Close()

	var content []byte
	tmp := make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		// 注意：检查是否读完一定要在判断出错之前
		if err == io.EOF {
			content = append(content, tmp[:n]...)
			fmt.Println("read file over")
			break
		} else if err != nil {
			fmt.Printf("read file %s failed, err:%v\n", filename, err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// ReadFile()：直接读取完整的文件
func ReadFile2(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file %s failed, err:%v\n", filename, err)
	}
	fmt.Println(string(content))
}

func main() {
	OpenFile1("./test.txt")
	ReadFile1("./test.txt")

	OpenFile2("./test1.txt")
	ReadFile2("./test1.txt")
}
