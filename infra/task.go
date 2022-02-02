package infra

import (
	"github.com/GenkiHirano/layered-architecture-practice/model"
	"github.com/GenkiHirano/layered-architecture-practice/repository"

	"github.com/jinzhu/gorm"
)

type taskRepository struct {
	Conn *gorm.DB
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{Conn: conn}
}

// Create taskの保存
func (tr *taskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}
