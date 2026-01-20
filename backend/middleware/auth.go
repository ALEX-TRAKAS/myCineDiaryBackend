package middleware

import (
	"mycinediarybackend/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET", ""))

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing auth token"})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid auth header format"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token claims"})
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid user_id in token"})
		}

		c.Set("userID", uint64(userIDFloat))
		return next(c)
	}
}
