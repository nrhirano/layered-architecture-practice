package config

import (
	"github.com/GenkiHirano/layered-architecture-practice/domain/model"

	"github.com/jinzhu/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(sample_db)/sample2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
