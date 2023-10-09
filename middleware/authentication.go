package middleware

import (
	"clean_arch/helpers"

	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt"
)

func Authentication() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(helpers.SECRET_KEY),
		SigningMethod: "HS256",
	})
}
