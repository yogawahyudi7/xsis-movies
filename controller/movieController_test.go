package controller_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"movies-xsis/common"
	"movies-xsis/constant"
	"movies-xsis/controller"
	mck "movies-xsis/mock"
	"movies-xsis/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllMovies_DataFound(t *testing.T) {

	expectedData := []model.Movie{
		{
			Id:    1,
			Title: "Pengabdi Setan",
		},
		{
			Id:    2,
			Title: "Spider-Man",
		},
	}

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(movieRepository)

	movieRepository.Mock.On("GetAll").Return(expectedData, expectedResponse)

	app := fiber.New()

	app.Get("/movies", movieController.GetAllMovie)

	req := httptest.NewRequest(http.MethodGet, "/movies", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("Response : ", httpResponse)

	assert.Equal(t, http.StatusOK, httpResponse.Code)

}

func TestGetAllMovies_DataNotFound(t *testing.T) {

	expectedData := []model.Movie{
		{
			Id:    1,
			Title: "Pengabdi Setan",
		},
		{
			Id:    2,
			Title: "Spider-Man",
		},
	}

	expectedResponse := common.StatusResponse{
		Code:  404,
		Error: errors.New(constant.ErrTestResponse),
	}

	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(movieRepository)

	movieRepository.Mock.On("GetAll").Return(expectedData, expectedResponse)

	app := fiber.New()

	app.Get("/movies", movieController.GetAllMovie)

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/movies", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusNotFound, httpResponse.Code)

}
