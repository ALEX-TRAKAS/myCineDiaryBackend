package handlers

import (
	"mycinediarybackend/models"
	"mycinediarybackend/services"
	"mycinediarybackend/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateThreadPost(c echo.Context) error {
	var threadPost models.ThreadPost
	if err := c.Bind(&threadPost); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	err := services.AddThreadPost(&threadPost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create thread post"})
	}
	return c.JSON(http.StatusCreated, threadPost)
}

func DeleteThreadPost(c echo.Context) error {
	threadPostID, err := utils.ParseUintParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid thread post ID"})
	}
	err = services.RemoveThreadPost(threadPostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete thread post"})
	}
	return c.NoContent(http.StatusNoContent)
}

func GetThreadPosts(c echo.Context) error {
	threadID, err := utils.ParseUintParam(c, "thread_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid thread ID"})
	}
	threadPosts, err := services.GetThreadPostsByThreadID(threadID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve thread posts"})
	}
	return c.JSON(http.StatusOK, threadPosts)
}

func UpdateThreadPost(c echo.Context) error {
	threadPostID, err := utils.ParseUintParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid thread post ID"})
	}
	var updateData struct {
		Body string `json:"body"`
	}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	err = services.UpdateThreadPostBody(threadPostID, updateData.Body, time.Now().UTC())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update thread post"})
	}
	return c.NoContent(http.StatusNoContent)
}
