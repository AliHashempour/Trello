package model

import "time"

const (
	RoleAdmin        = "Admin"
	RoleStandardUser = "User"
)

type UserWorkspaceRole struct {
	ID          uint   `gorm:"primarykey"`
	UserID      uint   `json:"user_id" gorm:"not null;"`
	WorkspaceID uint   `json:"workspace_id" gorm:"not null;"`
	Role        string `json:"role" gorm:"type:varchar(32);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (UserWorkspaceRole) TableName() string {
	return "user_workspace_roles"
}
