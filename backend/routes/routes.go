package routes

import (
	"mycinediarybackend/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.POST("/users", handlers.CreateUser)
	api.GET("/users/:id", handlers.GetUser)
	api.POST("/auth/register", handlers.Register)
	api.POST("/auth/login", handlers.Login)

}
