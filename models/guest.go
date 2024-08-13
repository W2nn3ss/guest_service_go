package models

import (
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required" validate:"required,string"`
	LastName  string `json:"last_name" binding:"required" validate:"required,string"`
	Email     string `json:"email" binding:"required,email" gorm:"unique" validate:"required,email"`
	Phone     string `json:"phone" binding:"required" gorm:"unique" validate:"required,phone"`
	Country   string `json:"country" validate:"required,string"`
}
