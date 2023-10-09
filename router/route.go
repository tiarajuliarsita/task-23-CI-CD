package router

import (
	"clean_arch/controller"
	"clean_arch/middleware"
	"clean_arch/repositories"
	"clean_arch/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(app *echo.Echo, db *gorm.DB) {

	repo := repositories.NewRepository(db)
	service := service.NewService(repo)
	controller := controller.NewController(service)
	r := app.Group("/users")
	r.POST("", controller.CreateUser)
	r.POST("/login", controller.LoginUser)

	auth := app.Group("/auth/users")
	auth.Use(middleware.Authentication())
	auth.GET("", controller.GetAllUsers)

}
