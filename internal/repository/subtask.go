package repository

import (
	"Trello/internal/model"
	"gorm.io/gorm"
)

type SubTaskRepository interface {
	GetBy(fields map[string]interface{}) (*model.SubTask, error)
	GetAll(taskID uint) ([]*model.SubTask, error)
	Create(subTask *model.SubTask) error
	Update(subTask *model.SubTask) error
	DeleteBy(fields map[string]interface{}) error
}

type subTaskRepository struct {
	db *gorm.DB
}

func NewSubTaskRepo(db *gorm.DB) SubTaskRepository {
	return &subTaskRepository{db}
}

func (r *subTaskRepository) GetBy(fields map[string]interface{}) (*model.SubTask, error) {
	var subTask model.SubTask
	err := r.db.Where(fields).First(&subTask).Error
	return &subTask, err
}

func (r *subTaskRepository) GetAll(taskID uint) ([]*model.SubTask, error) {
	var subTasks []*model.SubTask
	err := r.db.Where("task_id = ?", taskID).Find(&subTasks).Error
	return subTasks, err
}

func (r *subTaskRepository) Create(subTask *model.SubTask) error {
	return r.db.Create(subTask).Error
}

func (r *subTaskRepository) Update(subTask *model.SubTask) error {
	return r.db.Save(subTask).Error
}

func (r *subTaskRepository) DeleteBy(fields map[string]interface{}) error {
	return r.db.Where(fields).Delete(&model.SubTask{}).Error
}
