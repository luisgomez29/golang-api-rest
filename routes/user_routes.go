package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/controllers"
)

type UserRouter interface {
	UserRouters(e *echo.Group)
}

type userRoutes struct {
	userController controllers.UserController
}

func NewUserRoutes(userController controllers.UserController) UserRouter {
	return &userRoutes{userController}
}

func (ur *userRoutes) UserRouters(e *echo.Group) {
	e.GET("users", ur.userController.GetAll)
	e.GET("users/:id", ur.userController.FindById)
	e.POST("users", ur.userController.Create)
	e.PUT("users/:id", ur.userController.Update)
	e.DELETE("users/:id", ur.userController.Delete)
}
