package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"mycinediarybackend/models"
	"mycinediarybackend/services"

	"mycinediarybackend/repositories"
)

type MediaHandler struct{}

func NewMediaHandler() *MediaHandler {
	return &MediaHandler{}
}

func (h *MediaHandler) GetMediaReviews(c echo.Context) error {
	tmdbIDParam := c.Param("tmdb_id")
	tmdbID, err := strconv.Atoi(tmdbIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid tmdb_id"})
	}

	mediaType := c.QueryParam("media_type")
	if mediaType == "" {
		mediaType = "movie"
	}

	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	page := 1
	limit := 12

	if pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 {
			limit = l
		}
	}

	ctx := c.Request().Context()
	reviews, err := repositories.GetMediaReviewsPublic(ctx, tmdbID, mediaType, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, reviews)
}

func (h *MediaHandler) GetMediaByTMDBID(c echo.Context) error {
	ctx := c.Request().Context()

	tmdbID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "tmdb_id not found"})
	}

	mediaType := models.MediaType(c.QueryParam("media_type"))
	if mediaType != models.MediaTypeMovie && mediaType != models.MediaTypeTV {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "media_type must be 'movie' or 'tv'"})
	}

	media, err := services.GetMediaByTMDBID(ctx, tmdbID, mediaType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if media == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "media not found"})
	}

	return c.JSON(http.StatusOK, media)
}
