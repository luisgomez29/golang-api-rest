package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/auto"
	"github.com/luisgomez29/golang-api-rest/config"
	"github.com/luisgomez29/golang-api-rest/controllers"
	"github.com/luisgomez29/golang-api-rest/database"
	"github.com/luisgomez29/golang-api-rest/repositories"
	"github.com/luisgomez29/golang-api-rest/routes"
)

var resetTables = flag.Bool("rt", false, "Reset tables")

func main() {
	flag.Parse()
	config.Load()
	db := database.Connect()
	if db != nil {
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
	}

	if *resetTables {
		auto.Load(db)
	}

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)
	userRoutes := routes.NewUserRoutes(userController)

	productRepository := repositories.NewProductRepository(db)
	productController := controllers.NewProductController(productRepository, userRepository)
	productRouter := routes.NewProductRouter(productController)

	e := echo.New()
	apiV1 := e.Group("/api/v1/")
	routes.InitRoutes(apiV1, userRoutes, productRouter)
	e.Logger.Fatal(e.Start(config.PORT))
}