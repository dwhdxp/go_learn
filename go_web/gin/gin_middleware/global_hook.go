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

func hook1(c *gin.Context) {
	fmt.Println("hook1 in")
	c.Next()

	fmt.Println("hook1 out")
}

func main() {
	r := gin.Default()

	// 注册全局中间件
	r.Use(statCost(), hook1)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello dxp",
			"status":  http.StatusOK,
		})
	})

	r.Run(":9090")
}
