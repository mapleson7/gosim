package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Horses   []Horse `gorm:"foreignKey:Owner"`
}
