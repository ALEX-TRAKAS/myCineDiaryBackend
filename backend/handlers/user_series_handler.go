package handlers

import (
	"fmt"
	"mycinediarybackend/models"
	"mycinediarybackend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

func GetUserSeries(c echo.Context) error {
	userIDParam := c.Param("user_id")
	var userID uint64
	_, err := fmt.Sscanf(userIDParam, "%d", &userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid user ID",
		})
	}
	userSeriesList, err := services.GetUserSeries(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, userSeriesList)
}
