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
		GetAll(c echo.Context) error
		FindById(c echo.Context) error
		Create(c echo.Context) error
		Update(c echo.Context) error
		Delete(c echo.Context) error
	}

	userController struct {
		userRepository repositories.UserRepository
	}
)

func NewUserController(userRepository repositories.UserRepository) UserController {
	return &userController{userRepository}
}

func (ctl userController) GetAll(c echo.Context) error {
	users, err := ctl.userRepository.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (ctl *userController) FindById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}

	u, err := ctl.userRepository.FindById(uint32(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
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
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		return echo.ErrBadRequest
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	user, err := ctl.userRepository.Update(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (ctl *userController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}

	if err := ctl.userRepository.Delete(uint32(id)); err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, "user deleted successfully")
}
