package service

import (
	"clean_arch/entity"
	"clean_arch/model"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateUser(user entity.Core) (entity.Core, error) {
	args := m.Called(user)
	return args.Get(0).(entity.Core), args.Error(1)
}

func (m *MockService) FindAllUser() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockService) Login(email string, password string) (model.User, string, error) {
	args := m.Called(email, password)
	return args.Get(0).(model.User), args.String(1), args.Error(2)
}
