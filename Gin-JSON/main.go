package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type msg struct{
	Name string `json:"user"`
	Age int
}
func main() {
	r := gin.Default()

	r.GET("/json", func(context *gin.Context) {
		//方法1：使用map来进行序列化，
		/*data := map[string]interface{}{
			"name": "zhiqing",
			"message": "hello world",
			"age": 26,
		}*/
		//方法2 gin.H( )
		data := gin.H{"name": "zhiqing", "message": "hello world", "age": 26}
		context.JSON(http.StatusOK, data)
	})
	//方法3 结构体   灵活使用tag对结构体字段做定制化操作 就是json
	r.GET("/another_json", func(context *gin.Context) {
		data := msg{
			"zhiqing",
			26,
		}
		context.JSON(http.StatusOK, data)
	})

	r.Run("127.0.0.1:9090")
}
