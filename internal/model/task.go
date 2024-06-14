package model

import (
	"time"
)

type Task struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title" gorm:"type:varchar(255);not null"`
	Description   string    `json:"description" gorm:"type:text"`
	Status        string    `json:"status" gorm:"type:varchar(50)"`
	EstimatedTime float64   `json:"estimated_time"`
	ActualTime    float64   `json:"actual_time"`
	DueDate       time.Time `json:"due_date"`
	Priority      int       `json:"priority"`
	WorkspaceID   uint      `json:"workspace_id" gorm:"not null;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ImageURL      string    `json:"image_url" gorm:"type:text"`
}
