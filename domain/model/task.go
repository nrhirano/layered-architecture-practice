package model

import (
	"errors"
)

// Task taskの構造体
type Task struct {
	ID      int
	Title   string
	Content string
}

// NewTask taskのコンストラクタ
func NewTask(title, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	task := &Task{
		Title:   title,
		Content: content,
	}

	return task, nil
}
