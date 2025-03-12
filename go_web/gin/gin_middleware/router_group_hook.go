package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func statCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		fmt.Println("API cost:", cost)
	}
}

func main() {
	r := gin.Default()

	// 法一：为路由组注册中间件
	userGroup := r.Group("/ping/user", statCost())
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
				"path":   "/ping/user/index",
			})
		})
		userGroup.POST("/create", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
				"path":   "/ping/user/create",
			})
		})
	}

	// 法二：为路由组注册中间件
	shopGroup := r.Group("/ping/shop")
	shopGroup.Use(statCost())
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
				"path":   "/ping/shop/index",
			})
		})
		shopGroup.POST("/create", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
				"path":   "/ping/shop/create",
			})
		})
	}

	r.Run(":9090")
}
