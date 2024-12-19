/*
POST方法可以接收HTML表单提交的数据，gin中提供PostForm()、DefaultPostForm()等
来获取form表单提交的数据
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./login.html", "./index.html")
	router.GET("/ping/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// POST请求
	router.POST("/ping/login", func(c *gin.Context) {
		username := c.PostForm("username")
		// pssword := c.PostForm("password")
		msg := c.DefaultPostForm("msg", "welcome to web")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":    username,
			"Message": msg,
		})
	})

	router.Run(":9091")
}
