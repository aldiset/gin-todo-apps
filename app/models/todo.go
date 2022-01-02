package models

import (
	"time"
)

type ToDo struct {

	Id					uint		`json:"id" gorm:"primary_key"`
	ActivityGroupId		int			`json:"activity_group_id"`
	Title				string		`json:"title"`
	IsActive 			bool		`json:"is_active" gorm:"default:true"`
	Priority			string		`json:"priority" gorm:"default:'very-high'"`
	CreateAt			time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt 			time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP; OnUpdate:CURRENT_TIMESTAMP"`
	DeletedAt			*time.Time	`json:"deleted_at"`
}