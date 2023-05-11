package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//静态文件：html页面上用到的样式文件.css js文件 图片等等
func main() {
	//模板的定义
	r := gin.Default()
	//加载静态文件，要加在解析模板之前写上静态文件去哪里找
	r.Static("/static","src/practice/Gin_demo/statics")
	//模板的解析
	//可以写多个路径进行解析  解析了必须使用
	r.LoadHTMLFiles("src/practice/Gin_demo/templates/posts/index.tmpl","src/practice/Gin_demo/templates/users/index.tmpl")
	//也可以用下面的方法进行多个文件的解析
	r.LoadHTMLGlob("src/practice/Gin_demo/templates/**/*")

	//模板的渲染
	r.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "zhiqing",
		})
	})
	r.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "zhiqing",
		})
	})
	r.Run(":8080")//启动server
}
