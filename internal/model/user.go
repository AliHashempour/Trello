package model

import (
	"time"
)

type User struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Username   string       `json:"username"  gorm:"unique;not null"`
	Email      string       `json:"email"     gorm:"unique;not null"`
	Password   string       `json:"password"  gorm:"not null"`
	Workspaces []*Workspace `gorm:"many2many:user_workspace_roles;"`
}
