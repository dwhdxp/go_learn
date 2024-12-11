/*
模板库：text/template、html/template
模板使用：
1.定义模板文件：定义后缀为.tmpl或.tpl的文件(必须使用UTF8编码)，文件中使用{{ }}包裹和标识传入数据，
传入的数据可以通过 . 或 .FieldName访问，渲染时会将{{ }}内容进行替换，其他内容不变；
2.解析模板文件：通过 Parse()、ParseFiles()、ParseGlob()解析模板文件，得到模板对象(* Template)；
3.渲染模板文件：通过Execute()、ExecuteTemplate()渲染模板，即使用数据取替换模板文件中的{{ }}
*/
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//  1.定义模板

	// 2.解析模板
	t, err := template.ParseFiles("./tmpl_file/base.tmpl", "./tmpl_file/hello.tmpl", "./tmpl_file/ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
	}
	u1 := User{
		Name:   "dwh",
		Gender: "man",
		Age:    24,
	}
	m1 := map[string]interface{}{
		"name":   "dxp",
		"gender": "man",
		"age":    24,
	}
	// 3.渲染模板
	err = t.ExecuteTemplate(w, "hello.tmpl", map[string]interface{}{
		"m1": m1,
		"u1": u1,
	})
	if err != nil {
		fmt.Printf("execute template failed, err:%v\n", err)
	}
}

func main() {
	http.HandleFunc("/ping", sayHello)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Printf("http server start failed,err:%v\n", err)
		return
	}
}
