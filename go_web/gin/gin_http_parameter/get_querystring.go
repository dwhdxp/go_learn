/*
通常GET请求时，会在URL中使用querystring来传输参数，这些参数以键值对的形式附加在URL的末尾，通常由?开始，用&分隔
gin提供Query()、DefaultQuery()、GetQuery()等来获取querystring
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// name := c.Query("query") // 获取指定key的value
		// name := c.DefaultQuery("query", "lxz") // 获取指定value，若没有返回设置的默认值
		name, ok := c.GetQuery("query") // 返回布尔值来判断是否存在k-v
		if !ok {
			name = "nobody"
		}
		age := c.Query("age")
		msg := c.DefaultQuery("msg", "hello lxz")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
			"msg":  msg,
		})
	})

	router.Run(":9091")
}
