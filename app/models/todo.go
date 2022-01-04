package models

import (
	"time"
)

type Todo struct {

	Id					uint		`json:"id" gorm:"primary_key; index:idx;priority:1"`
	ActivityGroupId		int			`json:"activity_group_id" gorm:"index:idx"`
	Title				string		`json:"title" gorm:"index:idx"`
	IsActive 			bool		`json:"is_active" gorm:"default:true;index:idx"`
	Priority			string		`json:"priority" gorm:"default:'very-high';index:idx"`
	CreateAt			time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt 			time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP; OnUpdate:CURRENT_TIMESTAMP; index:idx"`
	DeletedAt			*time.Time	`json:"deleted_at" gorm:"index:idx"`
}