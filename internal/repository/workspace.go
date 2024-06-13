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
	var workspace model.Workspace
	if err := repo.db.First(&workspace, id).Error; err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (repo *WorkspaceRepo) GetAll() ([]*model.Workspace, error) {

	var workspace []*model.Workspace
	if err := repo.db.Find(&workspace).Error; err != nil {
		return nil, err
	}
	return workspace, nil
}

func (repo *WorkspaceRepo) Create(workspace *model.Workspace) error {
	return repo.db.Create(workspace).Error
}

func (repo *WorkspaceRepo) Update(workspace *model.Workspace) error {

	_, err := repo.GetByID(workspace.ID)
	if err != nil {
		return err
	}
	return repo.db.Save(workspace).Error
}

func (repo *WorkspaceRepo) Delete(id uint) error {
	_, err := repo.GetByID(id)
	if err != nil {
		return err
	}
	return repo.db.Delete(&model.Workspace{}, id).Error
}
