package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取URL路径参数 用：实现，返回的都是字符串类型 Param()
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	//下面冒号：后面的传过去，匹配的username赋值到网站路径上，匹配的address赋值到网站路径对应位置上
	r.GET("/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8080")
}
