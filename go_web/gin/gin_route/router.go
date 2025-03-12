package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Any内部定义了大多数请求方法，可以匹配多种请求
	r.Any("/ping", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"method": "PUT"})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
		}
	})

	// 处理所有未配置路由处理函数的路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "request failed!",
		})
	})

	r.Run(":9091")
}
