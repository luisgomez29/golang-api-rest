package main

import (
	"flag"
	"fmt"
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

	loginController := controllers.NewAuthController(db)
	loginRouter := routes.NewAuthRouter(loginController)

	e := echo.New()
	apiV1 := e.Group("/api/v1/")
	//apiV1.Use(middlewares.Authenticated())

	routes.InitRoutes(apiV1, userRoutes, productRouter, loginRouter)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.PORT)))
}
