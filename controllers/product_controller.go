package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/models"
	"github.com/luisgomez29/golang-api-rest/repositories"
	"net/http"
	"strconv"
)

type ProductController interface {
	All(echo.Context) error
	FindById(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type productController struct {
	productRepository repositories.ProductRepository
	userRepository    repositories.UserRepository
}

func NewProductController(pr repositories.ProductRepository, ur repositories.UserRepository) ProductController {
	return &productController{pr, ur}
}

func (ctl *productController) All(c echo.Context) error {
	p, err := ctl.productRepository.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, p)
}

func (ctl *productController) FindById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}
	p, err := ctl.productRepository.FindById(uint32(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, p)
}

func (ctl *productController) Create(c echo.Context) error {
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	u, err := ctl.userRepository.FindById(p.UserID)
	if err != nil {
		return err
	}
	p.User = *u
	product, err := ctl.productRepository.Create(p)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, product)
}

func (ctl *productController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}

	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return err
	}

	p, err = ctl.productRepository.Update(uint32(id), p)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, p)
}

func (ctl *productController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}
	rowsAffected, err := ctl.productRepository.Delete(uint32(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, rowsAffected)
}
