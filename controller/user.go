package controller

import (
	"clean_arch/entity"
	"clean_arch/model"
	"clean_arch/service"

	"github.com/labstack/echo/v4"
)

type controller struct {
	controller service.Service
}

func NewController(svc service.Service) *controller {
	return &controller{svc}
}

func (c *controller) GetAllUsers(e echo.Context) error {
	users, err := c.controller.FindAllUser()
	if err != nil {
		return e.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return e.JSON(200, echo.Map{
		"data": users,
	})
}

func (c *controller) CreateUser(e echo.Context) error {
	var userReq entity.UserRequest
	err := e.Bind(&userReq)
	if err != nil {
		return e.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	dataUser := entity.Core{
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	user, err := c.controller.CreateUser(dataUser)

	respondata := entity.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return e.JSON(200, echo.Map{
		"data":    respondata,
		"message": "succes create user",
	})
}

func (c *controller) LoginUser(e echo.Context) error {
	user := model.User{}
	e.Bind(&user)
	_, token, err := c.controller.Login(user.Email, user.Password)
	if err != nil {
		return e.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	return e.JSON(200, echo.Map{
		"token": token,
	})

}
