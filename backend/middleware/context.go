package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthGetUserID(c echo.Context) (uint64, error) {
	id, ok := c.Get("userID").(uint64)
	if !ok {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	return id, nil
}
