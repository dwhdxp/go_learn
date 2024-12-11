package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:9091/ping")
	if err != nil {
		fmt.Printf("resquest web failed, err:%v\n", err)
		return
	}
	// 程序结束后，必须关闭请求主体
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("resquest web response:%s\n", string(body))
}
