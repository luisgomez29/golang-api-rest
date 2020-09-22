package routes

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Group, routes *userRoutes) {
	routes.UserRouters(e)
	return
}
