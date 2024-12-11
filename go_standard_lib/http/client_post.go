package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	contentType := "application/json"
	data := `{"name":"dwh", "age":24}`
	resp, err := http.Post("http://localhost:9091/ping", contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}
