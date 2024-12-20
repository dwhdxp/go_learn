package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/ping/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	/*
		处理multipart/form-data提交文件时的默认内存限制为32MB
		可以通过r.MaxMultipartMemory来指定内存限制
	*/
	r.POST("ping/upload", func(c *gin.Context) {
		// 获取请求发送的文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
		// 将文件上传到指定路径
		// 1.拼接文件路径
		path := fmt.Sprintf("./%s", file.Filename)
		// 2.上传文件至指定路径
		c.SaveUploadedFile(file, path)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": fmt.Sprintf("%s upload successful", file.Filename),
		})
	})
	r.Run(":9091")
}
