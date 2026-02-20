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

func AddUserMovie(c echo.Context) error {
	ctx := c.Request().Context()

	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	type AddUserMovieRequest struct {
		TMDBMovieID int    `json:"tmdb_movie_id"`
		PosterPath  string `json:"poster_path"`
		Title       string `json:"title"`
	}

	var req AddUserMovieRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	if req.TMDBMovieID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "tmdb_movie_id is required",
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

	userMovie := models.UserMovie{
		UserID:      authUserID,
		TMDBMovieID: req.TMDBMovieID,
		PosterPath:  req.PosterPath,
		Title:       req.Title,
	}

	log.Printf(
		"AddUserMovie: user=%d tmdb_movie_id=%d poster_path=%s title=%s\n",
		authUserID,
		req.TMDBMovieID,
		req.PosterPath,
		req.Title,
	)

	if err := services.AddUserMovie(ctx, &userMovie); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, userMovie)
}

func RemoveUserMovie(c echo.Context) error {
	ctx := c.Request().Context()
	authUserID, err := middleware.AuthGetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	tmdbMovieID, err := strconv.Atoi(c.Param("tmdb_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid tmdb id",
		})
	}
	if err := services.RemoveUserMovie(ctx, authUserID, tmdbMovieID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User movie removed successfully",
	})
}

func GetUserMovies(c echo.Context) error {
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

	userMoviesList, err := services.GetUserMovies(ctx, authUserID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, userMoviesList)
}
