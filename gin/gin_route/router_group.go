package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 将具有公共URL前缀的路由划分为一个路由组
	pingGroup := r.Group("/ping")
	{
		pingGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "index"})
		})

		pingGroup.GET("/post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "post"})
		})
		
		// 嵌套路由组
		userGroup := pingGroup.Group("/user")
		{
			userGroup.GET("/login", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "user login"})
			})

			userGroup.GET("/regist", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "user regist"})
			})

			userGroup.GET("/index", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "user index"})
			})
		}
	}

	r.Run(":9091")
}
