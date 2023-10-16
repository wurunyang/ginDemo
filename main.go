package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建服务
	ginServer := gin.Default()

	// 创建接口
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello"})
	})

	// 启动服务
	ginServer.Run(":8082")
}
