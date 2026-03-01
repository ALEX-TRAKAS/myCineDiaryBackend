package routes

import (
	"mycinediarybackend/handlers"
	"mycinediarybackend/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")

	userHandler := handlers.NewUserHandler()
	authHandler := handlers.NewAuthHandler()
	mediaHandler := handlers.NewMediaHandler()

	// AUTH routes (public)
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
	api.POST("/refresh", authHandler.RefreshToken)
	api.GET("/media/:tmdb_id/reviews", mediaHandler.GetMediaReviews)

	// USER routes
	api.GET("/users/:id", userHandler.GetUser)

	// Protected routes (need JWT)
	auth := api.Group("/user")
	auth.Use(middleware.JWTMiddleware)

	auth.GET("", userHandler.GetCurrentUser)
	auth.POST("/logout", authHandler.Logout)
	auth.POST("/logout_all", authHandler.LogoutAll)

	auth.POST("/media", handlers.AddUserMedia)
	auth.GET("/media", handlers.GetUserMedia)
	auth.GET("/media/:tmdb_id", handlers.GetUserMediaByTMDBID)
	auth.DELETE("/media/:tmdb_id", handlers.RemoveUserMedia)

	auth.POST("/reviews", handlers.AddReview)
	auth.GET("/reviews", handlers.GetReviews)
	auth.DELETE("/reviews/:tmdb_id", handlers.RemoveReview)

	auth.POST("/threads", handlers.CreateThread)
	auth.POST("/threads/:id/posts", handlers.CreateThreadPost)
}
