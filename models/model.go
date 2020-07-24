package models

import (
	"time"
	database "wristband-nb-server/database"
)

var (
	db, error = database.Initdb()
)

// Model - 預設資料型態
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
