package handlers

import (
	"log"
	"mycinediarybackend/middleware"
	"mycinediarybackend/models"
	"mycinediarybackend/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddUserSeries(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	type AddUserSeriesRequest struct {
		TMDBSeriesID int    `json:"tmdb_series_id"`
		PosterPath   string `json:"poster_path"`
		Title        string `json:"title"`
	}

	var req AddUserSeriesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	if req.TMDBSeriesID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "tmdb_series_id is required",
		})
	}

	if req.PosterPath == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "poster_path is required",
		})
	}
	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "title is required",
		})
	}

	userSeries := models.UserSeries{
		UserID:       authUserID,
		TMDBSeriesID: req.TMDBSeriesID,
		PosterPath:   req.PosterPath,
		Title:        req.Title,
	}

	log.Printf(
		"AddUserSeries: user=%d tmdb_series_id=%d poster_path=%s title=%s\n",
		authUserID,
		req.TMDBSeriesID,
		req.PosterPath,
		req.Title,
	)

	if err := services.AddUserSeries(ctx, &userSeries); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, userSeries)
}

func RemoveUserSeries(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	tmdbSeriesID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid tmdb id"})
	}
	if err := services.RemoveUserSeries(ctx, authUserID, tmdbSeriesID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User series removed successfully"})
}

func GetUserSeries(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	userSeriesList, err := services.GetUserSeries(ctx, authUserID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, userSeriesList)
}
