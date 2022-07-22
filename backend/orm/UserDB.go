package orm

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	Email  string `json:"email" binding:"required" gorm:"not null"`
}