/*
HandlerFunc()：设置一个请求路径和请求处理器
ListenAndServe()：监听server地址并运行服务
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("key", "value")
	// 设置响应体
	w.Write([]byte("<h1>welcome my web</h1>"))
	fmt.Fprintf(w, "\nhello, %v\n", r.URL.Path)
}

func postHeadler(w http.ResponseWriter, r *http.Request) {
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	//r.ParseForm()
	//fmt.Println(r.PostForm)
	//fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	//defer r.Body.Close()
	// 2. 请求类型是application/json时从r.Body读取数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	// http.HandleFunc("/", getHandler)
	http.HandleFunc("/ping", getHandler)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v\n", err)
		return
	}
}
