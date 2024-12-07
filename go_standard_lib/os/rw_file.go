/*
1.读操作：file.Read([]byte)：读取到切片中；os.ReadFile(string)：直接读取文件全部内容；
2.写操作：file.Write([]byte)：写入切片数据；file.WriteString(string)：直接写入字符串数据；
os.WriteFile()：无需打开文件，直接向文件中写切片数据，以os.O_TRUNC；
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file %s error %v\n", path, err)
		return
	}
	defer file.Close()

	// file.Read()
	var content []byte
	tmp := make([]byte, 128)
	// 循环读
	for {
		n, err := file.Read(tmp)
		// 先判断是否读完，再判断是否出错，因为读完时err也不为nil;
		if err == io.EOF {
			// 要将当前读到的字节数也输出
			content = append(content, tmp[:n]...)
			break
		} else if err != nil {
			fmt.Printf("file.read file %s error %v\n", path, err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
	fmt.Println("file.read file over")

	// os.ReadFile()
	content, err = os.ReadFile(path)
	if err != nil {
		fmt.Printf("os.ReadFile %s error %v\n", path, err)
		return
	}
	fmt.Println(string(content))
	fmt.Println("os.readfile file over")
}

func writeFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer file.Close()

	// file.Write
	str := "you can be okey, please believe youself!"
	n, err := file.Write([]byte(str))
	if err != nil {
		fmt.Printf("file.write file failed, err:%v\n", err)
		return
	}
	fmt.Printf("file.write file size:%d\n", n)

	// file.WriteString
	n, err = file.WriteString(str)
	if err != nil {
		fmt.Printf("file.writestring file failed, err:%v\n", err)
		return
	}
	fmt.Printf("file.writestring file size:%d\n", n)

	// os.WriteString：以O_TRUNC模式写
	err = os.WriteFile(path, []byte(str), 0666)
	if err != nil {
		fmt.Printf("os.writefile failed, err:%v\n", err)
		return
	}
	fmt.Printf("os.writefile size:%d\n", n)
}

func main() {
	writeFile("./test.txt")
	readFile("./test.txt")
}
