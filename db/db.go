package db

import (
	"Score_System/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func InitDB(url string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Evaluation{}, &model.Item{}, &model.Level{}, &model.User{})
	if err != nil {
		return nil, err
	}
	return &DB{
		DB: db,
	}, nil
}
