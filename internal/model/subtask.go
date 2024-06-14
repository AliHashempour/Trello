package model

import (
	"time"
)

type SubTask struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TaskID      uint      `json:"task_id" gorm:"not null;"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	IsCompleted bool      `json:"is_completed" gorm:"type:boolean"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
