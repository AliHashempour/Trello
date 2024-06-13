package repository

import (
	"Trello/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uint) (*model.User, error)
	GetAll() ([]*model.User, error)
	Create(workspace *model.User) error
	Update(workspace *model.User) error
	Delete(id uint) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
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
	existingUser, err := r.GetByID(user.ID)
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

func (r *UserRepo) Delete(id uint) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}
	return r.db.Delete(&model.User{}, id).Error
}
