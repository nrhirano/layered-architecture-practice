package main

import (
	"github.com/GenkiHirano/layered-architecture-practice/interface/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	// taskRepository := infra.NewTaskRepository(config.NewDB())
	// taskService := service.NewTaskService(taskRepository)
	// taskHandler := handler.NewTaskHandler(taskService)

	// r := gin.New()
	// handler.InitRouting(r, taskHandler)
	// r.Logger.Fatal(r.Start(":8080"))

	r := gin.New()
	r.GET("/task", func(c *gin.Context) {
		handler.TaskHandler.Post()
	})
	r.PUT("/", func(c *gin.Context) {
		// ハンドラーを呼ぶ
	})
	r.POST("/", func(c *gin.Context) {
		// ハンドラーを呼ぶ
	})
	r.DELETE("/", func(c *gin.Context) {
		// ハンドラーを呼ぶ
	})

}
