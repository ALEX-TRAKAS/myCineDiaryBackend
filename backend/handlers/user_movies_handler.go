package handlers

import (
	"fmt"
	"net/http"

	"mycinediarybackend/models"
	"mycinediarybackend/services"

	"github.com/labstack/echo/v4"
)

func AddUserMovie(c echo.Context) error {
	var userMovie models.UserMovie
	if err := c.Bind(&userMovie); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}
	if err := services.AddUserMovie(&userMovie); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, userMovie)
}

func RemoveUserMovie(c echo.Context) error {
	var req struct {
		UserID      uint64 `json:"user_id"`
		TmdbMovieID int    `json:"tmdb_movie_id"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}
	if err := services.RemoveUserMovie(req.UserID, req.TmdbMovieID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User movie removed successfully",
	})
}

func GetUserMovies(c echo.Context) error {
	userIDParam := c.Param("user_id")
	var userID uint64
	_, err := fmt.Sscanf(userIDParam, "%d", &userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid user ID",
		})
	}
	userMoviesList, err := services.GetUserMovies(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, userMoviesList)
}
