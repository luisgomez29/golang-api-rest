package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/models"
	"github.com/luisgomez29/golang-api-rest/repositories"
	"net/http"
	"strconv"
)

type (
	UserController interface {
		Create(c echo.Context) error
		GetAll(c echo.Context) error
		Update(c echo.Context) error
	}

	userController struct {
		userRepository repositories.UserRepository
	}
)

func NewUserController(userRepository repositories.UserRepository) *userController {
	return &userController{userRepository}
}

func (ctl *userController) GetAll(c echo.Context) error {
	users, err := ctl.userRepository.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (ctl *userController) Create(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	user, err := ctl.userRepository.Create(user)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func (ctl *userController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	user, err = ctl.userRepository.Update(uint32(id), user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
