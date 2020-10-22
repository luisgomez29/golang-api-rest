package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisgomez29/golang-api-rest/config"
)

func Authenticated() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    config.SECRETKEY,
		SigningMethod: "HS512",
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/v1/login" {
				return true
			}
			return false
		},
	})
}
