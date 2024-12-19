/*
gin提供ShouldBind()等方法来快速获取请求中的各类参数(querystring、form表单、json等)到结构体中;
ShouldBind()会根据请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中各类参数;
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Name    string `form:"name" json:"name"`
	Age     int    `form:"age" json:"age"`
	Message string `form:"message" json:"message"`
}

func main() {
	router := gin.Default()

	// querystring eg:/ping?name=xxx&age=xxx&message=xxx
	router.GET("/ping", func(c *gin.Context) {
		var u1 UserInfo
		// 必须传入指针将参数绑定到结构体
		err := c.ShouldBind(&u1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("UserInfo:%v\n", u1)
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"name":    u1.Name,
				"age":     u1.Age,
				"message": u1.Message,
			})
		}
	})

	// json eg1:{"name":"xxx", "age":xxx, "message":"xxx"}、form eg2:name=xxx&age=xxx&message=xxx
	router.POST("/ping", func(c *gin.Context) {
		var u1 UserInfo
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		err := c.ShouldBind(&u1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("UserInfo:%v\n", u1)
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"name":    u1.Name,
				"age":     u1.Age,
				"message": u1.Message,
			})
		}
	})

	router.Run(":9091")
}
