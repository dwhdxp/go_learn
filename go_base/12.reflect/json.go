package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name"`
	Sex   string   `json:"sex"`
	Infos []string `json:"infos"`
}

func main() {
	user := User{"lxz", "female", []string{"phone", "beijing", "university"}}
	// 序列化：func json.Marshal(v any) ([]byte, error)
	jsonstr, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("json marshal succuss, jsonstr:%s\n", jsonstr)

	// 反序列化：func json.Unmarshal(data []byte, v any) error
	user1 := User{}
	err1 := json.Unmarshal(jsonstr, &user1)
	if err1 != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("json unmarshal success, user:%v\n", user1)
}
