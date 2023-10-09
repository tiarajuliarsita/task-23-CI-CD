package repositories

import (
	"clean_arch/entity"
	"clean_arch/model"
	"github.com/stretchr/testify/mock"
)

// MockRepository adalah struktur untuk mock repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateUser(user entity.Core) (entity.Core, error) {
	args := m.Called(user)
	return args.Get(0).(entity.Core), args.Error(1)
}

func (m *MockRepository) FindAllUser() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockRepository) Login(email string, password string) (model.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(model.User), args.Error(1)
}
