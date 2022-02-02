package handler

import (
	"github.com/GenkiHirano/layered-architecture-practice/handler"
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, handler handler.TaskHandler) {
	e.POST("/task", handler.Post())
}
