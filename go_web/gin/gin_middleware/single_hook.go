/*
1.gin允许在处理请求时，添加用户自定义的钩子(Hook)函数，即中间件，中间件适合处理一些公共的业务逻辑，
比如登录认证、权限校验、数据分页、记录日志、耗时统计等；
2.gin的中间件必须是gin.HandlerFunc类型；
3.gin中间件中使用goroutine时，不能使用原始的上下文，必须使用其只读副本c.Copy()
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func hook1(c *gin.Context) {
	fmt.Println("hook1 in")
	// 执行后续路由处理函数
	c.Next()
	// c.Abort() // 不执行后续的路由处理函数
	fmt.Println("hook1 out")
}

func hook2(c *gin.Context) {
	fmt.Println("hook2 in")
	// 使用goroutine时，不能使用原始的上下文，必须使用其只读副本c.Copy()
	go func(c *gin.Context) {
		fmt.Println("hook2 routine start")
		c.Set("message", "hello world!!!") // 修改的是请求上下文的拷贝，原始上下文
	}(c.Copy())
	c.Next()
	// c.Abort()
	fmt.Println("hook2 out")
}

// 统计接口耗时中间件
func statCost() gin.HandlerFunc {
	/*
		一般通过闭包来实现中间件，因为可以在返回HandlerFunc之前做一些准备工作
		例如连接数据库等
	*/
	return func(c *gin.Context) {
		// 可以通过Set()在请求上下文中设置值，后续路由处理函数通过Get()、MustGet()获取值
		start := time.Now()
		c.Set("message", "hello world")
		fmt.Println("cost in")
		c.Next()
		cost := time.Since(start)
		fmt.Println("API cost:", cost)
		fmt.Println("cost out")
	}
}

func main() {
	/*
		* gin.Default()默认使用了Logger, Recovery两个中间件
			* Logger将日志写入 gin.DefaultWriter，gin.DefaultWrite默认为os.Stdout；
			* Recovery会recover任何panic，如果有panic的话，会写入500响应码；
		* gin.New()可以创建一个无默认中间件的Engine，即gin.Default == gin.New + Logger + Recovery
	*/
	r := gin.Default()
	// handlers ...HandlerFunc，可接收多个HandlerFunc参数
	r.GET("/ping", hook1, hook2, statCost(), func(c *gin.Context) { // type HandlerFunc func(*Context)
		fmt.Println("ping in")
		msg := c.MustGet("message").(string) // 获取之前中间件中设置的值
		fmt.Println(msg)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": msg,
		})
	})

	r.Run(":9090")
}
