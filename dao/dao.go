package dao

import (
	"Score_System/db"
)

type DAO struct {
	DB *db.DB
}

func InitDao(db *db.DB) *DAO {
	return &DAO{DB: db}
}
