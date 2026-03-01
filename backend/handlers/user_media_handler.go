package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"mycinediarybackend/middleware"
	"mycinediarybackend/models"
	"mycinediarybackend/services"

	"github.com/labstack/echo/v4"
)

func AddUserMedia(c echo.Context) error {
	ctx := c.Request().Context()

	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	type AddUserMediaRequest struct {
		TMDBID       int                `json:"tmdb_id"`
		MediaType    models.MediaType   `json:"media_type"`
		Title        string             `json:"title"`
		PosterPath   string             `json:"poster_path"`
		BackdropPath string             `json:"backdrop_path"`
		Overview     string             `json:"overview"`
		ReleaseDate  *time.Time         `json:"release_date"`
		Rating       *int               `json:"rating,omitempty"`
		Progress     *int               `json:"progress,omitempty"`
		Notes        string             `json:"notes,omitempty"`
		IsFavorite   bool               `json:"is_favorite,omitempty"`
		Status       models.MediaStatus `json:"status,omitempty"`
	}

	var req AddUserMediaRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if req.TMDBID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "tmdb_id is required"})
	}

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "title is required"})
	}

	if req.MediaType != models.MediaTypeMovie && req.MediaType != models.MediaTypeTV {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "media_type must be 'movie' or 'tv'"})
	}

	if req.Status == "" {
		req.Status = models.StatusWatchlist
	}

	userMedia := models.UserMedia{
		UserID:       authUserID,
		TMDBID:       req.TMDBID,
		MediaType:    req.MediaType,
		Title:        req.Title,
		PosterPath:   req.PosterPath,
		BackdropPath: req.BackdropPath,
		Overview:     req.Overview,
		ReleaseDate:  req.ReleaseDate,
		Rating:       req.Rating,
		Progress:     req.Progress,
		Notes:        req.Notes,
		IsFavorite:   req.IsFavorite,
		Status:       req.Status,
	}

	log.Printf(
		"AddUserMedia: user=%d tmdb_id=%d type=%s status=%s\n",
		authUserID,
		req.TMDBID,
		req.MediaType,
		req.Status,
	)

	if err := services.AddUserMedia(ctx, &userMedia); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, userMedia)
}

func RemoveUserMedia(c echo.Context) error {
	ctx := c.Request().Context()

	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	tmdbID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid tmdb_id"})
	}

	mediaType := models.MediaType(c.QueryParam("media_type"))
	if mediaType != models.MediaTypeMovie && mediaType != models.MediaTypeTV {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "media_type must be 'movie' or 'tv'"})
	}

	if err := services.RemoveUserMedia(ctx, authUserID, tmdbID, mediaType); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User media removed successfully"})
}

func GetUserMedia(c echo.Context) error {
	ctx := c.Request().Context()

	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 12
	}

	mediaType := models.MediaType(c.QueryParam("media_type"))
	if mediaType != "" && mediaType != models.MediaTypeMovie && mediaType != models.MediaTypeTV {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid media_type filter"})
	}

	userMediaList, err := services.GetUserMedia(ctx, authUserID, page, limit, mediaType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, userMediaList)
}

func GetUserMediaByTMDBID(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	tmdbID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "tmdb_id not found"})
	}
	mediaType := models.MediaType(c.QueryParam("media_type"))
	if mediaType != models.MediaTypeMovie && mediaType != models.MediaTypeTV {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "media_type must be 'movie' or 'tv'"})
	}
	userMedia, err := services.GetUserMediaByTMDBID(ctx, authUserID, tmdbID, mediaType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if userMedia == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user media not found"})
	}
	return c.JSON(http.StatusOK, userMedia)
}
