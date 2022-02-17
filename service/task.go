package service

import (
	"github.com/GenkiHirano/layered-architecture-practice/domain/model"
	"github.com/GenkiHirano/layered-architecture-practice/repository"
)

// TaskService task serviceのinterface
type TaskService interface {
	Create(title, content string) (*model.Task, error)
	Get(id int) (*model.Task, error)
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

// Create taskを保存するときのユースケース
func (tu *taskService) Get(id int) (*model.Task, error) {
	task, err := model.NewTask2(id)
	if err != nil {
		return nil, err
	}

	getTask, err := tu.taskRepo.Get(task)
	if err != nil {
		return nil, err
	}

	return getTask, nil
}
