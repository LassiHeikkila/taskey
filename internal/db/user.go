package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"unique,not null"`
}
