package models

import (
	"time"
)

type Activity struct {

	Id			uint		`json:"id" gorm:"primary_key,index"`
	Email		string		`json:"email" gorm:"index"`
	Title		string		`json:"title" gorm:"index"`
	CreateAt	time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP;index"`
	UpdateAt 	time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP; OnUpdate:CURRENT_TIMESTAMP;index"`
	DeletedAt	*time.Time	`json:"deleted_at" gorm:"index"`
}