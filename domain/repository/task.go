package repository

import (
	"github.com/GenkiHirano/layered-architecture-practice/domain/model"
)

// TaskRepository task repositoryのinterface
type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
}
