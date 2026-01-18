package handlers

import (
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

func AddUserSeries(c echo.Context) error {
	var userSeries models.UserSeries
	if err := c.Bind(&userSeries); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}
	if err := services.AddUserSeries(&userSeries); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, userSeries)
}

func RemoveUserSeries(c echo.Context) error {
	var req struct {
		UserID       uint64 `json:"user_id"`
		TmdbSeriesID int    `json:"tmdb_series_id"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}
	if err := services.RemoveUserSeries(req.UserID, req.TmdbSeriesID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User series removed successfully",
	})
}
