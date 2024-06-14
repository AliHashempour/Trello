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

	var user model.User
	err := r.db.First(&user, userWorkspaceRole.UserID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	var workspace model.Workspace
	err = r.db.First(&workspace, userWorkspaceRole.WorkspaceID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("workspace not found")
		}
		return err
	}

	return r.db.Create(userWorkspaceRole).Error
}

func (r *userWorkspaceRepositoryImpl) GetUsersByWorkspaceID(workspaceID uint) ([]*model.User, error) {
	var userWorkspaceRoles []model.UserWorkspaceRole
	err := r.db.Where("workspace_id = ?", workspaceID).Find(&userWorkspaceRoles).Error
	if err != nil {
		return nil, err
	}

	if len(userWorkspaceRoles) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	userIDs := make([]uint, len(userWorkspaceRoles))
	for i, uwr := range userWorkspaceRoles {
		userIDs[i] = uwr.UserID
	}

	var users []*model.User
	err = r.db.Where("id IN ?", userIDs).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userWorkspaceRepositoryImpl) UpdateRole(userWorkspaceRole *model.UserWorkspaceRole) error {
	if !isValidRole(userWorkspaceRole.Role) {
		return errors.New("invalid role")
	}

	var existing model.UserWorkspaceRole
	err := r.db.Where("user_id = ? AND workspace_id = ?", userWorkspaceRole.UserID, userWorkspaceRole.WorkspaceID).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return r.db.Model(&existing).Update("role", userWorkspaceRole.Role).Error
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
