package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"mycinediarybackend/services"
)

func GetUserActivity(c echo.Context) error {

	userID := c.Param("id")

	activities, err := services.GetUserActivity(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch activities",
		})
	}

	return c.JSON(http.StatusOK, activities)
}
