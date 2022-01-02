package models

import (
	"time"
)

type Activity struct {

	Id			uint		`json:"id" gorm:"primary_key"`
	Email		string		`json:"email"`
	Title		string		`json:"title"`
	CreateAt	time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt 	time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP; OnUpdate:CURRENT_TIMESTAMP"`
	DeletedAt	*time.Time	`json:"deleted_at"`
}