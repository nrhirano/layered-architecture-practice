package handler

import (
	"net/http"

	"github.com/GenkiHirano/layered-architecture-practice/service"
	"github.com/gin-gonic/gin"
)

// TaskHandler task handlerのinterface
type TaskHandler interface {
	Post() gin.HandlerFunc
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

		// serviceのCreateへジャンプします
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
// func (th *taskHandler) Get() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi((c.Param("id")))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
// 		}

// 		foundTask, err := th.taskService.Get(id)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
// 		}

// 		res := responseTask{
// 			ID:      foundTask.ID,
// 			Title:   foundTask.Title,
// 			Content: foundTask.Content,
// 		}

// 		c.JSON(http.StatusOK, res)
// 	}
// }
