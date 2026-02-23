package handlers

import (
	"log"
	"net/http"
	"strconv"

	"mycinediarybackend/middleware"
	"mycinediarybackend/models"
	"mycinediarybackend/services"

	"github.com/labstack/echo/v4"
)

func AddReview(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	var req models.Review
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if req.TMDBID == 0 || req.MediaType == "" || req.Rating < 1 || req.Rating > 10 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "tmdb_id, media_type, and rating(1-10) required"})
	}

	req.UserID = userID
	log.Printf("AddReview: user=%d tmdb=%d type=%s rating=%d\n", userID, req.TMDBID, req.MediaType, req.Rating)

	if err := services.AddReview(ctx, &req); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, req)
}

func RemoveReview(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	tmdbID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tmdb_id"})
	}

	mediaType := c.QueryParam("media_type")
	if mediaType == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "media_type required"})
	}

	if err := services.RemoveReview(ctx, userID, tmdbID, mediaType); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "review removed"})
}

func GetReviews(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	mediaType := c.QueryParam("media_type")

	reviews, err := services.GetReviews(ctx, userID, page, limit, mediaType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, reviews)
}
