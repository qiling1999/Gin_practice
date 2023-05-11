package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query string参数
		//方法2通过DefaultQuery获取请求种携带的query string 参数，取不到就用默认值
		username := c.DefaultQuery("username", "zhiqing")
		//方法2通过Query获取请求种携带的query string 参数
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}