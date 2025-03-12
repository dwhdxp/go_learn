/*
1.加载静态文件(.css, .js)加载：Static()，注意在html中加上相对路径；
2.模板解析：LoadHTMLFiles()和LoadHTMLGlob()；
3.模板渲染：返回HTML数据
*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 加载静态文件：相对路径/self表示路径./statics
	r.Static("/self", "./statics")

	// 模板解析
	// r.LoadHTMLGlob("tmpl_file/**/*",) // **表示目录，*表示文件
	r.LoadHTMLFiles("tmpl_file/indexs/index.html", "tmpl_file/posts/index.html", "tmpl_file/other.html")

	r.GET("/ping/indexs", func(c *gin.Context) {
		// 注意：如果没有在html中定义name，那么name就是文件名
		c.HTML(http.StatusOK, "indexs/index.html", gin.H{
			"title": "hi, index",
		})
	})

	r.GET("/ping/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "hi, posts",
		})
	})

	// 返回下载的前端模板：1.解析html文件；2.更改html文件中的静态文件路径为/self；
	r.GET("ping/other", func(c *gin.Context) {
		c.HTML(http.StatusOK, "other.html", nil)
	})

	r.Run(":9091")
}
