/*
通过gin的Param()获取URI的path参数，注意多个URI的匹配不要冲突
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/ping/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"Name": name,
			"Age":  age,
		})
	})

	router.Run(":9091")
}
