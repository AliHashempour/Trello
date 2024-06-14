package model

import (
	"time"
)

type Workspace struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string               `json:"name"        gorm:"type:varchar(32);not null"`
	Description    string               `json:"description" gorm:"type:text;"`
	UserWorkspaces []*UserWorkspaceRole `gorm:"foreignKey:WorkspaceID;constraint:OnDelete:CASCADE;"`
	Tasks          []*Task              `json:"tasks" gorm:"foreignKey:WorkspaceID;constraint:OnDelete:CASCADE;"`
}
