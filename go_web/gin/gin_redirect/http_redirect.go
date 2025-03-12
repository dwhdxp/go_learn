package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 外部重定向
	r.GET("/ping", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	// 内部重定向
	r.GET("/ping/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ping")
	})

	r.Run(":9091")
}
