package main

import (
	"mycinediarybackend/config"
	"mycinediarybackend/database"
	"mycinediarybackend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load config
	config.Load()

	// Connect to database
	database.Connect(config.GetEnv("DATABASE_URL", ""))

	// Create Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	routes.RegisterRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
