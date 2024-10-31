package router

import (
	"github.com/gin-gonic/gin"
)

func newRouter(opts ...gin.OptionFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.(error).Error(),
		})
	}))
	return engine.With(opts...)
}
