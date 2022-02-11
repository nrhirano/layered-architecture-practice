package repository

import (
	"github.com/GenkiHirano/layered-architecture-practice/domain/model"
)

// TaskRepository task repositoryのinterface
type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	Get(id int) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	// Delete(task *model.Task) error
}
