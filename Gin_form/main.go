package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取form表单提交的参数
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	//解析页面
	r.LoadHTMLFiles("src/practice/Gin_demo/Gin_form/login.html", "src/practice/Gin_demo/Gin_form/index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//访问Login post的请求
	//获取form表单提交的参数，方法一
	r.POST("/login", func(c *gin.Context) {
		//获取form表单提交的参数，方法一
		username := c.PostForm("username")
		password := c.PostForm("password")
		// DefaultPostForm取不到值时会返回指定的默认值  方法2
		//username := c.DefaultPostForm("username", "zhiqing")
		//password := c.DefaultPostForm("password", "519")
		c.HTML(http.StatusOK,"Gin_form/index.html", gin.H{
			"Name": username,
			"PWD": password,
		})
		//输出json结果给调用方
		//c.JSON(http.StatusOK, gin.H{
		//	"message":  "ok",
		//	"username": username,
		//	"password":  password,
		//})
	})
	r.Run(":8080")
}