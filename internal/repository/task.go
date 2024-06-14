package repository

import (
	"Trello/internal/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetBy(fields map[string]interface{}) (*model.Task, error)
	GetAll(workspaceID uint) ([]*model.Task, error)
	Create(task *model.Task) error
	Update(task *model.Task) error
	DeleteBy(fields map[string]interface{}) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetBy(fields map[string]interface{}) (*model.Task, error) {
	var task model.Task
	err := r.db.Where(fields).First(&task).Error
	return &task, err
}

func (r *taskRepository) GetAll(workspaceID uint) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.Where("workspace_id = ?", workspaceID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteBy(fields map[string]interface{}) error {
	return r.db.Where(fields).Delete(&model.Task{}).Error
}
