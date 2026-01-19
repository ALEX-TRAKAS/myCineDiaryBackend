package handlers

import (
	"mycinediarybackend/middleware"
	"mycinediarybackend/models"
	"mycinediarybackend/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddUserSeries(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	var userSeries models.UserSeries
	if err := c.Bind(&userSeries); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	userSeries.UserID = authUserID
	if err := services.AddUserSeries(ctx, &userSeries); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, userSeries)
}

func RemoveUserSeries(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	tmdbSeriesID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid tmdb id"})
	}
	if err := services.RemoveUserSeries(ctx, authUserID, tmdbSeriesID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User series removed successfully"})
}

func GetUserSeries(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userSeriesList, err := services.GetUserSeries(ctx, authUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, userSeriesList)
}
