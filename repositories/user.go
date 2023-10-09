package repositories

import (
	"clean_arch/entity"
	"clean_arch/model"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	CreateUser(user entity.Core) (entity.Core, error)
	FindAllUser() ([]model.User, error)
	Login(email string, password string) (model.User, error)
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB}
}

func (r *repository) CreateUser(user entity.Core) (entity.Core, error) {
	userData := model.User{
		Email:    user.Email,
		Password: user.Password,
	}
	err := r.DB.Create(&userData).Error
	if err != nil {
		return user, err
	}
	user.ID = userData.ID
	user.Email = userData.Email
	user.CreatedAt = userData.CreatedAt
	user.UpdatedAt = userData.UpdatedAt
	user.Password = userData.Password
	return user, nil
}

func (r *repository) FindAllUser() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) Login(email string, password string) (model.User, error) {
	user := model.User{}
	err := r.DB.Where("password = ? AND email = ?", password, email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
