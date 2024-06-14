package repository

import (
	"Trello/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetBy(fields map[string]interface{}) (*model.User, error)
	GetAll() ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	DeleteBy(fields map[string]interface{}) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetBy(fields map[string]interface{}) (*model.User, error) {
	var user model.User

	err := r.db.Where(fields).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) Update(user *model.User) error {
	existingUser, err := r.GetBy(map[string]interface{}{"id": user.ID})
	if err != nil {
		return err
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}

	if user.Password != "" {
		existingUser.Password = user.Password
	}

	err = r.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(existingUser).Error
	if err != nil {
		return err
	}
	*user = *existingUser
	return nil
}

func (r *UserRepo) DeleteBy(fields map[string]interface{}) error {
	var user model.User
	userRecord, err := r.GetBy(fields)
	if err != nil {
		return err
	}

	err = r.db.Where("user_id = ?", userRecord.ID).Delete(&model.UserWorkspaceRole{}).Error
	if err != nil {
		return err
	}

	return r.db.Where(fields).Delete(&user).Error
}
