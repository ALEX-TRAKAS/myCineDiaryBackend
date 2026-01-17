package handlers

import (
	"mycinediarybackend/services"

	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid input",
		})
	}

	if len(req.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Password too short",
		})
	}

	err := services.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User created",
	})
}

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request",
		})
	}

	user, err := services.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Invalid credentials",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
