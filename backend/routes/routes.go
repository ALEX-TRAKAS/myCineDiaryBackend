package routes

import (
	"mycinediarybackend/handlers"
	"mycinediarybackend/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")

	userHandler := handlers.NewUserHandler()

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/users/:id", userHandler.GetUser)

	auth := api.Group("/user")
	auth.Use(middleware.JWTMiddleware)

	auth.GET("", userHandler.GetCurrentUser)

	auth.POST("/movies", handlers.AddUserMovie)
	auth.GET("/movies", handlers.GetUserMovies)
	auth.DELETE("/movies/:tmdb_id", handlers.RemoveUserMovie)
	auth.POST("/series", handlers.AddUserSeries)
	auth.GET("/series", handlers.GetUserSeries)
	auth.DELETE("/series/:tmdb_id", handlers.RemoveUserSeries)
	auth.POST("/threads", handlers.CreateThread)
	auth.POST("/threads/:id/posts", handlers.CreateThreadPost)
}
