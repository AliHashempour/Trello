package model

import (
	"time"
)

const (
	RoleAdmin        = "Admin"
	RoleStandardUser = "Standard User"
)

type UserWorkspaceRole struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	WorkspaceID uint   `gorm:"not null"`
	Role        string `json:"role" gorm:"type:varchar(32);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (UserWorkspaceRole) TableName() string {
	return "user_workspace_roles"
}