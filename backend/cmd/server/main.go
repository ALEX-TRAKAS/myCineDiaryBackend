package main

import (
	"context"
	"mycinediarybackend/config"
	"mycinediarybackend/database"
	"mycinediarybackend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.Load()

	database.Connect()
	defer database.DB.Close(context.Background())

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
