package repository

import (
	"log"

	"github.com/GenkiHirano/layered-architecture-practice/domain/model"
	"xorm.io/xorm"
)

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	Get(id int) (*model.Task, error)
}

type taskRepository struct {
	Engine *xorm.Engine
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(engine *xorm.Engine) TaskRepository {
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

// Create taskの保存
func (tr *taskRepository) Get(id int) (*model.Task, error) {
	task := &model.Task{
		ID: id,
	}

	_, err := tr.Engine.Where("id = ?", task.ID).Get(task)
	if err != nil {
		log.Fatal(err)
	}

	return task, nil
}
