package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/controllers"
)

type (
	UserRouter interface {
		UserRouters(e *echo.Group)
	}

	userRoutes struct {
		userController controllers.UserController
	}
)

func NewUserRoutes(userController controllers.UserController) *userRoutes {
	return &userRoutes{userController}
}

func (ur *userRoutes) UserRouters(e *echo.Group) {
	e.GET("/users", ur.userController.GetAll)
	e.POST("/users", ur.userController.Create)
	e.PUT("/users/:id", ur.userController.Update)
}
