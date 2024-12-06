/*
使用接口实现一个简易日志库，要求可以向文件中写，也可以向终端写
*/
package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(string)
}

// 向文件中写日志
type FileLogger struct {
	FileName string
}

func (fl *FileLogger) Info(msg string) {
	var file *os.File
	var err error
	file, err = os.OpenFile(fl.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error, err:%v\n", err)
		return
	}
	defer file.Close()
	msg += "\n"
	n, err1 := file.WriteString(msg)
	if err1 != nil {
		fmt.Printf("write file error, err:%v\n", err1)
	}
	fmt.Printf("write %d bytes to file\n", n)
}

// 向终端写日志
type TerminalLogger struct{}

func (tl *TerminalLogger) Info(msg string) {
	fmt.Println(msg)
}

func writeLog(msg string) {
	var logger Logger
	logger = &FileLogger{"log.txt"}
	logger.Info(msg)

	logger = &TerminalLogger{}
	logger.Info(msg)
}

func main() {
	str := "hello world, next to meet you"
	writeLog(str)
}
