package infra

import (
	"log"

	"github.com/GenkiHirano/layered-architecture-practice/domain/model"
	"github.com/GenkiHirano/layered-architecture-practice/domain/repository"
	"xorm.io/xorm"
)

type taskRepository struct {
	Engine *xorm.Engine
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(engine *xorm.Engine) repository.TaskRepository {
	return &taskRepository{Engine: engine}
}

// Create taskの保存
func (tr *taskRepository) Create(task *model.Task) (*model.Task, error) {
	newTask := &model.Task{
		ID:      task.ID,
		Title:   task.Title,
		Content: task.Content,
	}

	_, err := tr.Engine.Table("task").Insert(newTask)
	if err != nil {
		log.Fatal(err)
	}

	return newTask, nil
}
