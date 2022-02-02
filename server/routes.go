package handler

import (
	"github.com/GenkiHirano/layered-architecture-practice/handler"
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, taskHandler handler.TaskHandler) {

	e.POST("/task", taskHandler.Post())
	e.GET("/task/:id", taskHandler.Get())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())

}
