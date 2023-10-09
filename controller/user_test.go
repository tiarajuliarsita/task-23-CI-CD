package controller

import (
	"clean_arch/model"
	"clean_arch/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	mockService := new(service.MockService)
	ctrl := NewController(mockService)

	expectedUsers := []model.User{
		{Email: "user1@example.com", Password: "password1"},
		{Email: "user2@example.com", Password: "password2"},
	}

	mockService.On("FindAllUser").Return(expectedUsers, nil)

	// Setup Echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Panggil GetAllUsers di controller
	err := ctrl.GetAllUsers(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
