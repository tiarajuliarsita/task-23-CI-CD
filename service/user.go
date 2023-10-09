package service

import (
	"clean_arch/entity"
	"clean_arch/helpers"
	"clean_arch/model"
	"clean_arch/repositories"
)

type Service interface {
	CreateUser(entity.Core) (entity.Core, error)
	FindAllUser() ([]model.User, error)
	Login(email string, password string) (model.User, string, error)
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) *service {
	return &service{repo}
}

func (s *service) CreateUser(user entity.Core) (entity.Core, error) {
	data,err := s.repo.CreateUser(user)
	if err != nil {
		return data,err
	}
	return data,nil
}

func (s *service) FindAllUser() ([]model.User, error) {
	// var user model.User
	users, err := s.repo.FindAllUser()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *service) Login(email string, password string) (model.User, string, error) {
	user, err := s.repo.Login(email, password)
	if err != nil {
		return user, "", err
	}
	token := helpers.GenerateToken(email, int(user.ID))
	return user, token, nil
}
