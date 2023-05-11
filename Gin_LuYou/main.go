package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由和路由组
//路由就是请求的路径和请求的方法，处理逻辑
func main() {
	r := gin.Default()
	//下面就是GET请求的路由，
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	//此外，还有一个可以匹配所有请求方法的Any方法如下
	r.Any("/any", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET",})
		case http.MethodPost: //这里是别人定义好的常量
			c.JSON(http.StatusOK, gin.H{"method": "POST",})
		}
	})
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"method": "GET",})})
		userGroup.GET("/login", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"method": "GET",})})
		userGroup.POST("/login", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"method": "POST",})})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"method": "GET",})})
		shopGroup.GET("/cart", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"method": "GET",})})
		shopGroup.POST("/checkout", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"method": "POST",})})
	}
	//路由组支持嵌套
	//shopGroup := r.Group("/shop")
	//	{
	//		shopGroup.GET("/index", func(c *gin.Context) {...})
	//		shopGroup.GET("/cart", func(c *gin.Context) {...})
	//		shopGroup.POST("/checkout", func(c *gin.Context) {...})
	//		// 嵌套路由组
	//		xx := shopGroup.Group("xx")
	//		xx.GET("/oo", func(c *gin.Context) {...})
	//	}
	//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回views/404.html页面。
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	r.Run()
}