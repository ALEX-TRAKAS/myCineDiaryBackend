package utils

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseUintParam(c echo.Context, name string) (uint64, error) {
	param := c.Param(name)
	if param == "" {
		return 0, fmt.Errorf("missing param: %s", name)
	}

	value, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}
