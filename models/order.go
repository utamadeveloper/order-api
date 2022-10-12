package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `json:"user_id" gorm:"not null" binding:"required"`
	Item   string `json:"item" gorm:"not null" binding:"required"`
}
