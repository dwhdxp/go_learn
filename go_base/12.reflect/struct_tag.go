package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `info:"name" doc:"我的名字"`
	Age  int    `info:"age" doc:"我的年龄"`
}

func printStructTag(x interface{}) {
	t := reflect.TypeOf(x).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Printf("taginfo=%s tagdoc=%s\n", taginfo, tagdoc)
	}
}

func main() {
	user := User{Name: "dxp", Age: 18}
	printStructTag(&user)
}
