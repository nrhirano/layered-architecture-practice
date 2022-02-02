package handler

import (
	"github.com/labstack/echo"
)

// TaskHandler task handlerのinterface
type TaskHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
