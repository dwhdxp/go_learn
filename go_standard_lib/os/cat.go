// 模拟实现linux下的cat命令
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //注意是字符
		if err == io.EOF {
			// 退出之前将已读到的内容输出
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	// 解析命令行参数
	flag.Parse()
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
