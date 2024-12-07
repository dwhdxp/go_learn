package main

import (
	"fmt"
	"io"
	"os"
)

// 实现拷贝文件函数
func CopyFile(dstFile, srcFile string) (n int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open srcfile failed, err:%v\n", err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open dstfile failed, err:%v\n", err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	n, err := CopyFile("./test1.txt", "./test.txt")
	if err != nil {
		fmt.Printf("copyfile failed, err:%v\n", err)
		return
	}
	fmt.Printf("copy file size:%d\n", n)
}
