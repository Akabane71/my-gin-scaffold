package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/logger"
)

// 注册路由

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	return r
}
