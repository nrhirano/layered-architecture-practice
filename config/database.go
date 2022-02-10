package config

import (
	"log"

	"github.com/GenkiHirano/layered-architecture-practice/domain/model"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// NewDB DBと接続する
func NewDB() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:password@tcp(db:3306)/test_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	engine.Sync2(model.Task{})

	return engine
}
