package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/controllers"
)

type AuthRouter interface {
	AuthRoutes(e *echo.Group)
}

type authRouter struct {
	authController controllers.AuthController
}

func NewAuthRouter(lc controllers.AuthController) AuthRouter {
	return &authRouter{lc}
}

func (lr *authRouter) AuthRoutes(e *echo.Group) {
	e.POST("login", lr.authController.Login)
	e.POST("verify_token", lr.authController.VerifyToken)
	e.POST("refresh_token", lr.authController.RefreshToken)
}
