package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	Detail string `json:"detail"`
	UserID uint   `json:"-"`
}
