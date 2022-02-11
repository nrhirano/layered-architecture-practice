package handler

import (
	"github.com/gin-gonic/gin"
)

// InitRouting routesの初期化
func InitRouting(c *gin.Engine, taskHandler TaskHandler) {

	c.POST("/task", taskHandler.Post())
	c.GET("/task/:id", taskHandler.Get())
	c.PUT("/task/:id", taskHandler.Put())
	// c.DELETE("/task/:id", taskHandler.Delete())
}
