package handlers

import (
	"net/http"
	"strconv"

	"mycinediarybackend/middleware"
	"mycinediarybackend/models"
	"mycinediarybackend/services"

	"github.com/labstack/echo/v4"
)

func AddUserMovie(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	var userMovie models.UserMovie
	if err := c.Bind(&userMovie); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}
	userMovie.UserID = authUserID
	if err := services.AddUserMovie(ctx, &userMovie); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, userMovie)
}

func RemoveUserMovie(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	tmdbMovieID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid tmdb id",
		})
	}
	if err := services.RemoveUserMovie(ctx, authUserID, tmdbMovieID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User movie removed successfully",
	})
}

func GetUserMovies(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userMoviesList, err := services.GetUserMovies(ctx, authUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, userMoviesList)
}
