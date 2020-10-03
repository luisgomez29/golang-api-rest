package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/controllers"
)

type ProductRouter interface {
	ProductRoutes(e *echo.Group)
}

type productRouter struct {
	productController controllers.ProductController
}

func NewProductRouter(pc controllers.ProductController) ProductRouter {
	return &productRouter{pc}
}

func (pr *productRouter) ProductRoutes(e *echo.Group) {
	e.GET("products", pr.productController.All)
	e.GET("products/:id", pr.productController.FindById)
	e.POST("products", pr.productController.Create)
	e.PUT("products/:id", pr.productController.Update)
	e.DELETE("products/:id", pr.productController.Delete)
}
