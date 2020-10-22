package routes

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Group, ur UserRouter, pr ProductRouter, lr AuthRouter) {
	ur.UserRouters(e)
	pr.ProductRoutes(e)
	lr.AuthRoutes(e)
	return
}
