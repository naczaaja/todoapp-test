package orm

import "gorm.io/gorm"

type TodoDB struct {
	gorm.Model
	Message  string `json:"message"`
	CreateBy int    `json:"createBy" gorm:"not null"`
}