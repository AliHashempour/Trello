package repository

import (
	"Trello/internal/model"
	"errors"
	"gorm.io/gorm"
)

type UserWorkspace interface {
	Create(userWorkspaceRole *model.UserWorkspaceRole) error
	GetUsersByWorkspaceID(workspaceID uint) ([]*model.User, error)
	UpdateRole(userWorkspaceRole *model.UserWorkspaceRole) error
	Delete(userID, workspaceID uint) error
}

type userWorkspaceRepositoryImpl struct {
	db *gorm.DB
}

func NewUserWorkspaceRepository(db *gorm.DB) UserWorkspace {
	return &userWorkspaceRepositoryImpl{db: db}
}

func (r *userWorkspaceRepositoryImpl) Create(userWorkspaceRole *model.UserWorkspaceRole) error {
	if !isValidRole(userWorkspaceRole.Role) {
		return errors.New("invalid role")
	}

	return r.db.Create(userWorkspaceRole).Error
}

func (r *userWorkspaceRepositoryImpl) GetUsersByWorkspaceID(workspaceID uint) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Model(&model.Workspace{Model: gorm.Model{ID: workspaceID}}).
		Association("Users").Find(&users)
	return users, err
}

func (r *userWorkspaceRepositoryImpl) UpdateRole(userWorkspaceRole *model.UserWorkspaceRole) error {
	if !isValidRole(userWorkspaceRole.Role) {
		return errors.New("invalid role")
	}

	return r.db.Model(&model.UserWorkspaceRole{}).
		Where("user_id = ? AND workspace_id = ?", userWorkspaceRole.UserID, userWorkspaceRole.WorkspaceID).
		Update("role", userWorkspaceRole.Role).Error
}

func (r *userWorkspaceRepositoryImpl) Delete(userID, workspaceID uint) error {
	return r.db.Where("user_id = ? AND workspace_id = ?", userID, workspaceID).
		Delete(&model.UserWorkspaceRole{}).Error
}

func isValidRole(role string) bool {
	validRoles := []string{model.RoleAdmin, model.RoleStandardUser}
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}