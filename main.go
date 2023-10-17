package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 定义一个中间键，可以在请求到达业务逻辑之间做一些预处理
func myHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("session-id", 11111)
		c.Next()
	}
}

func main() {
	// 创建服务
	ginServer := gin.Default()
	// 注册中间件，如果没有在特定的接口中指定使用中间件，则中间件对所有接口生效。否则只针对特定接口生效
	ginServer.Use(myHandler())

	// 创建接口 GET方法
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello"})
	})

	// 创建接口 POST方法
	ginServer.POST("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "you post a user"})
	})

	// 获取查询字符串中的参数，使用Query方法实现
	ginServer.GET("/user/info", func(c *gin.Context) {
		log.Println("============>", c.MustGet("session-id"))
		id := c.Query("id")
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
	})

	// RESTful API 获取参数
	ginServer.GET("/users/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
	})

	// 获取请求体中的参数
	ginServer.POST("/users", func(c *gin.Context) {
		body, _ := c.GetRawData()
		var res map[string]any
		json.Unmarshal(body, &res)
		c.JSON(http.StatusOK, res)
	})

	// 路由重定向（如果不生效，可能是浏览器缓存的问题，清理下就OK了）
	ginServer.GET("/runyangwu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://github.com/wurunyang/GolandLearningProject")
	})

	// 路由组，用于统一管理一组请求。这里简单写了，并没有实现什么功能
	testGroup := ginServer.Group("/test")
	{
		// 路由地址为 /test/cases
		testGroup.GET("/cases")
		testGroup.POST("/cases")
		testGroup.PUT("/cases/:id")
		testGroup.DELETE("/cases/:id")
	}

	// 启动服务
	ginServer.Run(":8082")
}
