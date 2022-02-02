package main

import (
	"github.com/GenkiHirano/layered-architecture-practice/config"
	"github.com/GenkiHirano/layered-architecture-practice/handler"
	"github.com/GenkiHirano/layered-architecture-practice/infra"
	"github.com/GenkiHirano/layered-architecture-practice/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	taskRepository := infra.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := echo.New()
	handler.InitRouting(e, taskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
