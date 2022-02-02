package handler

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, handler handler.TaskHandler) {

	e.POST("/task", handler.Post())
	e.GET("/task/:id", handler.Get())
	e.PUT("/task/:id", handler.Put())
	e.DELETE("/task/:id", handler.Delete())

}
