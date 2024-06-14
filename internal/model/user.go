package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string       `json:"username"  gorm:"unique;not null"`
	Email      string       `json:"email"     gorm:"unique;not null"`
	Password   string       `json:"password"  gorm:"not null"`
	Workspaces []*Workspace `gorm:"many2many:user_workspace_roles;"`
}
