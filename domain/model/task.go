package model

import (
	"errors"
)

// Task taskの構造体
type Task struct {
	ID      int    `xorm:"id pk autoincr"`
	Title   string `xorm:"title"`
	Content string `xorm:"content"`
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

// Set taskのセッター
func (t *Task) Set(title, content string) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	t.Title = title
	t.Content = content

	return nil
}
