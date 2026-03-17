package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"mycinediarybackend/middleware"
	"mycinediarybackend/services"
)

func GetUserActivity(c echo.Context) error {
	ctx := c.Request().Context()

	userID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
	}

	activities, err := services.GetUserActivity(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch activities",
		})
	}

	return c.JSON(http.StatusOK, activities)
}
