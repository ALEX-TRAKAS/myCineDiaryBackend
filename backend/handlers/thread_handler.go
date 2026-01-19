package handlers

import (
	"mycinediarybackend/models"
	"mycinediarybackend/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllThreads(c echo.Context) error {
	threads, err := services.GetAllThreads()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, threads)
}

func CreateThread(c echo.Context) error {
	var thread models.Thread
	if err := c.Bind(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}
	thread.CreatedAt = time.Now()
	if err := services.CreateThread(thread); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, thread)
}

func GetThreadByID(c echo.Context) error {
	threadID := c.Param("id")
	thread, err := services.GetThreadByID(threadID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, thread)
}

func DeleteThread(c echo.Context) error {
	threadID := c.Param("id")
	if err := services.DeleteThread(threadID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Thread deleted successfully",
	})
}
