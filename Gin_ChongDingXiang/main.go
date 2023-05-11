package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//请求的重定向，分为临时重定向和永久重定向
//一个请求访问到我的服务器后，把他转给别的服务器

func main() {

	r := gin.Default()
	//重定向
	r.GET("/test0", func(c *gin.Context) {
		//正常访问
		//c.JSON(http.StatusOK,gin.H{
		//	"static": "ok",
		//})
		//HTTP重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	//路由重定向
	r.GET("/test1", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})


	r.Run()
}
