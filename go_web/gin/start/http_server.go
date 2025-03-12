/*
使用gin开启一个http server：
1.创建路由引擎：gin.Default()
2.定义方法并设置请求路径和处理器：
3.监听并开启端口：Run()
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建一个默认路由引擎
	r := gin.Default()

	// gin支持RESTful风格 API的开发
	// 2.设置请求路径和请求处理器
	r.GET("/ping", func(c *gin.Context) {
		// c.JSON返回json风格数据
		c.JSON(http.StatusOK, gin.H{ // H == map[string]interface{}
			"message": "method GET",
		})
	})

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "method POST",
		})
	})

	r.PUT("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "method PUT",
		})
	})

	r.DELETE("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "method DELETE",
		})
	})

	// 3.启动HTTP服务，默认在8080端口
	// r.Run()
	r.Run(":9091")
}
