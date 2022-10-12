package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	FirstName string  `json:"first_name" gorm:"not null" binding:"required"`
	LastName  string  `json:"last_Name" gorm:"not null" binding:"required"`
	Email     string  `json:"email" gorm:"not null" binding:"required"`
	Orders    []Order `json:"orders" gorm:"foreignkey:UserID"`
}
