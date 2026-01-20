package handlers

import (
	"net/http"
	"strconv"

	"mycinediarybackend/middleware"
	"mycinediarybackend/models"
	"mycinediarybackend/services"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	user, err := services.GetUserByID(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	publicUser := models.PublicUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(http.StatusOK, publicUser)
}

func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := middleware.AuthGetUserID(c)

	user, err := services.GetUserByID(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	publicUser := models.PublicUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(http.StatusOK, publicUser)
}
