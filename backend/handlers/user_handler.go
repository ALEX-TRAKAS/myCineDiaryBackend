package handlers

import (
	"net/http"
	"strconv"

	"mycinediarybackend/middleware"
	"mycinediarybackend/models"
	"mycinediarybackend/services"
	"mycinediarybackend/utils"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if err := services.Register(ctx, req.Username, req.Email, req.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "registration successful"})
}

func (h *UserHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	user, err := services.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
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
