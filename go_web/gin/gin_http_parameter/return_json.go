package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 方法一：拼接json，gin.H是map[string]interface{}缩写
	//data := map[string]interface{}{
	//	"Name":    "dxp",
	//	"Message": "hello lxz",
	//	"Age":     18,
	//}
	data := gin.H{
		"Name":    "dxp",
		"Message": "hello lxz",
		"Age":     18,
	}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, data) // json序列化
	})

	// 方法二：使用结构体
	router.GET("/ping/other", func(c *gin.Context) {
		type msg struct {
			Name    string `json:"name"` // 使用tag标签灵活显式字段名
			Message string
			Age     int
		}
		var data msg
		data.Name = "dxp"
		data.Message = "hello lxz"
		data.Age = 24
		c.JSON(http.StatusOK, data)
	})

	router.Run(":9091")
}
