package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rhavyx/react_go_watch-movies/backend_golang/src/models"
)

func (app *application) getOneMovie(c *gin.Context) {
	movie_id := c.Param("movie_id")
	fmt.Println(movie_id)

	_id, err := strconv.Atoi(movie_id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJson(c, err)
		return
	}

	movie := models.Movie{
		ID:          _id,
		Title:       "Some title",
		Description: "Some description",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	c.JSON(http.StatusOK, movie)

}

func (app *application) getAllMovies(c *gin.Context) {

}
