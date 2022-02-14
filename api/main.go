package main

import (
	"github.com/GenkiHirano/layered-architecture-practice/config"
	"github.com/GenkiHirano/layered-architecture-practice/interface/handler"
	"github.com/GenkiHirano/layered-architecture-practice/repository"
	"github.com/GenkiHirano/layered-architecture-practice/service"
	"github.com/gin-gonic/gin"
)

func main() {
	taskRepository := repository.NewTaskRepository(config.NewDB())
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	r := gin.New()
	handler.InitRouting(r, taskHandler)
	r.Run()
}
