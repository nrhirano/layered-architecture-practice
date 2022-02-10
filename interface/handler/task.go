package handler

import (
	"net/http"
	"strconv"

	"github.com/GenkiHirano/layered-architecture-practice/service"
	"github.com/gin-gonic/gin"
)

// TaskHandler task handlerのinterface
type TaskHandler interface {
	Post() gin.HandlerFunc
	Get() gin.HandlerFunc
	Put() gin.HandlerFunc
	// Delete() gin.HandlerFunc
}

type taskHandler struct {
	taskService service.TaskService
}

// NewTaskHandler task handlerのコンストラクタ
func NewTaskHandler(taskService service.TaskService) TaskHandler {
	return &taskHandler{taskService: taskService}
}

type requestTask struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTask struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Post taskを保存するときのハンドラー
func (th *taskHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestTask
		if err := c.Bind(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		createdTask, err := th.taskService.Create(req.Title, req.Content)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      createdTask.ID,
			Title:   createdTask.Title,
			Content: createdTask.Content,
		}

		c.JSON(http.StatusCreated, res)
	}
}

// Get taskを取得するときのハンドラー
func (th *taskHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		foundTask, err := th.taskService.FindByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      foundTask.ID,
			Title:   foundTask.Title,
			Content: foundTask.Content,
		}

		c.JSON(http.StatusOK, res)
	}
}

// Put taskを更新するときのハンドラー
func (th *taskHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		var req requestTask
		if err := c.Bind(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		updatedTask, err := th.taskService.Update(id, req.Title, req.Content)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      updatedTask.ID,
			Title:   updatedTask.Title,
			Content: updatedTask.Content,
		}

		c.JSON(http.StatusOK, res)
	}
}

// // Delete taskを削除するときのハンドラー
// func (th *taskHandler) Delete() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
// 		}

// 		err = th.taskService.Delete(id)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
// 		}

// 		// c.NoContent(http.StatusNoContent)
// 	}
// }
