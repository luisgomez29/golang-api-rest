package main

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/auto"
	"github.com/luisgomez29/golang-api-rest/config"
	"github.com/luisgomez29/golang-api-rest/controllers"
	"github.com/luisgomez29/golang-api-rest/database"
	"github.com/luisgomez29/golang-api-rest/repositories"
	"github.com/luisgomez29/golang-api-rest/routes"
)

func main() {
	config.Load()

	db := database.Connect()
	if db != nil {
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
	}

	auto.Load(db)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)
	userRoutes := routes.NewUserRoutes(userController)

	e := echo.New()
	apiV1 := e.Group("/api/v1.0")
	routes.InitRoutes(apiV1, userRoutes)
	e.Logger.Fatal(e.Start(config.PORT))
}
