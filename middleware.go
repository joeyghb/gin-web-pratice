package main

import (
	"github.com/gin-gonic/gin"
)

func middleware1(c *gin.Context) {

	c.JSON(200, gin.H{
        "mw1": "exec middleware1",
    })
	// 执行该中间件之后的逻辑
	c.Next()
}

func middleware2(c *gin.Context) {

	c.JSON(200, gin.H{
        "mw2_start": "arrive at middleware2",
	})
	// 执行该中间件之前，先跳到流程的下一个方法
	c.Next()
	// 流程中的其他逻辑已经执行完了
	c.JSON(200, gin.H{
        "mw2_end": "exec middleware2",
	})
}

func main() {
	router := gin.Default()
	// 注册一个路由，使用了 middleware1，middleware2 两个中间件
	router.GET("/someGet", middleware1, middleware2, handler)

	router.Run(":18080")
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
        "handler": "exec handler",
	})
}
