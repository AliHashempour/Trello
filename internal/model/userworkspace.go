package model

import "gorm.io/gorm"

const (
	RoleAdmin        = "Admin"
	RoleStandardUser = "User"
)

type UserWorkspaceRole struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"not null;"`
	WorkspaceID uint   `json:"workspace_id" gorm:"not null;"`
	Role        string `json:"role" gorm:"type:varchar(32);not null"`
}

func (UserWorkspaceRole) TableName() string {
	return "user_workspace_roles"
}
