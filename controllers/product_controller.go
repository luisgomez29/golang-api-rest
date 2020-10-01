package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/repositories"
	"net/http"
)

type ProductController interface {
	All(c echo.Context) error
}

type productController struct {
	productRepository repositories.ProductRepository
}

func NewProductController(pr repositories.ProductRepository) ProductController {
	return &productController{pr}
}

func (ctl *productController) All(c echo.Context) error {
	u, err := ctl.productRepository.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
