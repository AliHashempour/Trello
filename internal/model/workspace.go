package model

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	Name        string `json:"name"        gorm:"type:varchar(32);not null"`
	Description string `json:"description" gorm:"type:text;"`
}