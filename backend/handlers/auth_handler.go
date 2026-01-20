package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
	"mycinediarybackend/services"
	"mycinediarybackend/utils"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	_, err := services.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "registration successful"})
}

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	ip := c.RealIP()
	deviceID := c.Request().Header.Get("X-Device-ID")
	if deviceID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "missing device id",
		})
	}

	user, err := services.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	accessToken, _ := utils.GenerateJWT(user.ID)
	familyID := uuid.New().String()
	refreshToken, err := services.CreateRefreshToken(ctx, user.ID, familyID, ip, deviceID)

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api/auth/refresh",
	})

	return c.JSON(http.StatusOK, echo.Map{"access_token": accessToken})
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing refresh token"})
	}

	token := cookie.Value

	currentIP := c.RealIP()

	stored, err := services.ValidateRefreshToken(ctx, token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid refresh token"})
	}

	if stored.LastUsedIP != "" && stored.LastUsedIP != currentIP {
		_ = repositories.RevokeFamilyTokensByFamilyID(ctx, stored.FamilyID)

		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "session compromised, please login again",
		})
	}
	_ = repositories.RevokeRefreshTokenByID(ctx, stored.ID)

	newToken, err := services.CreateRefreshToken(
		ctx,
		stored.UserID,
		stored.FamilyID,
		stored.CreatedIP,
		stored.DeviceID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to rotate refresh token"})
	}

	_ = repositories.UpdateRefreshTokenUsage(ctx, stored.ID, currentIP)

	accessToken, _ := utils.GenerateJWT(stored.UserID)

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    newToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api/auth/refresh",
	})

	return c.JSON(http.StatusOK, echo.Map{"access_token": accessToken})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "missing refresh token"})
	}

	tokenSHA := utils.SHA256(cookie.Value)
	rt, err := repositories.GetRefreshTokenBySHA(ctx, tokenSHA)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid refresh token"})
	}

	_ = repositories.RevokeRefreshTokenByID(ctx, rt.ID)

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
		Path:     "/api/auth/refresh",
	})

	return c.JSON(http.StatusOK, echo.Map{"message": "logged out"})
}

func (h *AuthHandler) LogoutAll(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "missing refresh token"})
	}

	tokenSHA := utils.SHA256(cookie.Value)
	rt, err := repositories.GetRefreshTokenBySHA(ctx, tokenSHA)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid refresh token"})
	}

	_ = repositories.RevokeFamilyTokensByFamilyID(ctx, rt.FamilyID)

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
		Path:     "/api/auth/refresh",
	})

	return c.JSON(http.StatusOK, echo.Map{"message": "logged out from all devices"})
}
