package service

import (
	"github.com/GenkiHirano/layered-architecture-practice/domain/model"
	"github.com/GenkiHirano/layered-architecture-practice/domain/repository"
)

// TaskService task serviceのinterface
type TaskService interface {
	Create(title, content string) (*model.Task, error)
}

type taskService struct {
	taskRepo repository.TaskRepository
}

// NewTaskService task serviceのコンストラクタ
func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

// Create taskを保存するときのユースケース
func (tu *taskService) Create(title, content string) (*model.Task, error) {
	task, err := model.NewTask(title, content)
	if err != nil {
		return nil, err
	}

	createdTask, err := tu.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
