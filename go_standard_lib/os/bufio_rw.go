/*
bufio是在file的基础上封装了一层API，操作时先将数据放入缓冲区，然后再操作；
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadByBufIO(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			break
		} else if err != nil {
			fmt.Printf("read file by bufio failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	//if err = scanner.Err(); err != nil {
	//	fmt.Printf("read file err:%v\n", err)
	//}
}

func WriteByBufIO(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 2; i++ {
		// 将数据写入buf中
		writer.WriteString("believe yourself, everything is going to be okey!")
	}
	// 刷新buf到文件
	writer.Flush()
}

func main() {
	WriteByBufIO("./test1.txt")
	ReadByBufIO("./test1.txt")
}
