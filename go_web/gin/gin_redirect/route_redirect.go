package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping/a", func(c *gin.Context) {
		// 跳转到重定向地址/ping/b对应的路由处理函数
		c.Request.URL.Path = "/ping/b"
		r.HandleContext(c)
	})

	r.GET("/ping/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	})

	r.Run(":9091")

}
