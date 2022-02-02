package usecase

import (
	"github.com/GenkiHirano/layered-architecture-practice/model"
	"github.com/GenkiHirano/layered-architecture-practice/repository"
)

// TaskUsecase task usecaseのinterface
type TaskUsecase interface {
	Create(title, content string) (*model.Task, error)
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

// NewTaskUsecase task usecaseのコンストラクタ
func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

// Create taskを保存するときのユースケース
func (tu *taskUsecase) Create(title, content string) (*model.Task, error) {
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
