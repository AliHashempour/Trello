package repository

import (
	"Trello/internal/model"
	"gorm.io/gorm"
)

type Workspace interface {
	GetByID(id uint) (*model.Workspace, error)
	GetAll() ([]*model.Workspace, error)
	Create(workspace *model.Workspace) error
	Update(workspace *model.Workspace) error
	Delete(id uint) error
}
type WorkspaceRepo struct {
	db *gorm.DB
}

func NewWorkspaceRepo(db *gorm.DB) *WorkspaceRepo {
	return &WorkspaceRepo{db: db}
}

func (repo *WorkspaceRepo) GetByID(id uint) (*model.Workspace, error) {
	return nil, nil
}

func (repo *WorkspaceRepo) GetAll() ([]*model.Workspace, error) {
	return nil, nil
}
func (repo *WorkspaceRepo) Create(workspace *model.Workspace) error {
	return nil
}
func (repo *WorkspaceRepo) Update(workspace *model.Workspace) error {
	return nil
}
func (repo *WorkspaceRepo) Delete(id uint) error {
	return nil
}
