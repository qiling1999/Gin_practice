package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。
//这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，
//比如登录认证、权限校验、数据分页、记录日志、耗时统计等。

//Gin中的中间件必须是一个gin.HandlerFunc类型。
//例如我们像下面的代码一样定义一个统计请求耗时的中间件。

//定义中间件
// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "zhiqing") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		fmt.Printf("cost%v\n", cost)
	}
}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}


func main() {
	r := gin.Default()

	// 注册一个全局中间件
	r.Use(StatCost(), )

	//或者给/index路由单独注册中间件如StatCost()（可注册多个）
	r.GET("/index", StatCost(), indexHandler)

	r.GET("/test1", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		indexHandler(c)
	})
	//为路由组注册中间件  写法1
	shopGroup := r.Group("/shop", StatCost())
	{
		shopGroup.GET("/index", indexHandler)
	}
	//为路由组注册中间件  写法2
	shopGroup2 := r.Group("/shop")
	shopGroup2.Use(StatCost())
	{
		shopGroup2.GET("/index", indexHandler)
	}

	//中间件注意事项
	//gin默认中间件
	//gin.Default()默认使用了Logger和Recovery中间件，其中：
	//
	//Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	//Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	//如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。
	//
	//gin中间件中使用goroutine
	//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。

	r.Run()
}
