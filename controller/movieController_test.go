package controller_test

import (
	"bytes"
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

	"github.com/go-playground/validator/v10"
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

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("GetAll").Return(expectedData, expectedResponse)

	app := fiber.New()

	app.Get("/movies", movieController.GetAllMovies)

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

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("GetAll").Return(expectedData, expectedResponse)

	app := fiber.New()

	app.Get("/movies", movieController.GetAllMovies)

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/movies", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusNotFound, httpResponse.Code)

}

func TestAddMovie_DataSaved(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.AddMovieRequest{
		Title:       "Power Rangers",
		Description: "Pahlawan super pembela kebeneran",
		Rating:      8,
		Image:       "",
	}
	movieRepository.Mock.On("Add", request).Return(expectedResponse)

	app := fiber.New()

	app.Post("/movies", movieController.AddMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusOK, httpResponse.Code)

}

func TestAddMovie_ServerUnderMaintenance(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  500,
		Error: errors.New(constant.ErrTestResponse),
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.AddMovieRequest{
		Title:       "Power Rangers",
		Description: "Pahlawan super pembela kebeneran",
		Rating:      8,
		Image:       "",
	}
	movieRepository.Mock.On("Add", request).Return(expectedResponse)

	app := fiber.New()

	app.Post("/movies", movieController.AddMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusInternalServerError, httpResponse.Code)

}

func TestAddMovie_ValidationFailed(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.AddMovieRequest{
		Title:       "Power Rangers",
		Description: "",
		Rating:      8,
		Image:       "",
	}
	movieRepository.Mock.On("Add", request).Return(expectedResponse)

	app := fiber.New()

	app.Post("/movies", movieController.AddMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

}

func TestAddMovie_InvalidJSONParameters(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.AddMovieRequest{
		Title:       "Power Rangers",
		Description: "Pahlawan super pembela kebeneran",
		Rating:      8,
		Image:       "",
	}
	movieRepository.Mock.On("Add", request).Return(expectedResponse)

	app := fiber.New()

	app.Post("/movies", movieController.AddMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(requestBody))

	// comment the code below to make apicall cant read the json request propertly
	// req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

}

func TestGetMovie_DataFound(t *testing.T) {

	expectedData := model.Movie{
		Id:    1,
		Title: "Pengabdi Setan",
	}

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("GetById", 1).Return(expectedData, expectedResponse)

	app := fiber.New()

	app.Get("/movies/:id", movieController.GetMovie)

	req := httptest.NewRequest(http.MethodGet, "/movies/1", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("Response : ", httpResponse)

	assert.Equal(t, http.StatusOK, httpResponse.Code)

}

func TestGetMovie_InvalidRouteParamters(t *testing.T) {

	expectedData := model.Movie{
		Id:    1,
		Title: "Pengabdi Setan",
	}

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("GetById", 1).Return(expectedData, expectedResponse)

	app := fiber.New()

	app.Get("/movies/:id", movieController.GetMovie)

	req := httptest.NewRequest(http.MethodGet, "/movies/a", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("Response : ", httpResponse)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

}

func TestGetMovie_DataNotFound(t *testing.T) {

	expectedData := model.Movie{
		Id:    1,
		Title: "Pengabdi Setan",
	}

	expectedResponse := common.StatusResponse{
		Code:  404,
		Error: errors.New(constant.ErrTestResponse),
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("GetById", 1).Return(expectedData, expectedResponse)

	app := fiber.New()

	app.Get("/movies/:id", movieController.GetMovie)

	req := httptest.NewRequest(http.MethodGet, "/movies/1", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("Response : ", httpResponse)

	assert.Equal(t, http.StatusNotFound, httpResponse.Code)

}

func TestDeleteMovie_DataDeleted(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("Delete", 1).Return(expectedResponse)

	app := fiber.New()

	app.Delete("/movies/:id", movieController.DeleteMovie)

	req := httptest.NewRequest(http.MethodDelete, "/movies/1", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("Response : ", httpResponse)

	assert.Equal(t, http.StatusOK, httpResponse.Code)

}

func TestDeleteMovie_InvalidRouteParamters(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("Delete", 1).Return(expectedResponse)

	app := fiber.New()

	app.Delete("/movies/:id", movieController.DeleteMovie)

	req := httptest.NewRequest(http.MethodDelete, "/movies/a", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("Response : ", httpResponse)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

}

func TestDeleteMovie_ServerUnderMaintenance(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  500,
		Error: errors.New(constant.ErrTestResponse),
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	movieRepository.Mock.On("Delete", 1).Return(expectedResponse)

	app := fiber.New()

	app.Delete("/movies/:id", movieController.DeleteMovie)

	req := httptest.NewRequest(http.MethodDelete, "/movies/1", nil)

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("Response : ", httpResponse)

	assert.Equal(t, http.StatusInternalServerError, httpResponse.Code)

}

func TestUpdateMovie_DataUpdated(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  200,
		Error: nil,
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.UpdateMovieRequest{
		Id:          1,
		Title:       "Power Rangers",
		Description: "Pahlawan super pembela kebeneran",
		Rating:      8,
		Image:       "",
	}
	movieRepository.Mock.On("Update", request).Return(expectedResponse)

	app := fiber.New()

	app.Patch("/movies", movieController.UpdateMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPatch, "/movies", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusOK, httpResponse.Code)

}

func TestUpdateMovie_ServerUnderMaintenance(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  500,
		Error: errors.New(constant.ErrTestResponse),
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.UpdateMovieRequest{
		Id:          1,
		Title:       "Power Rangers",
		Description: "Pahlawan super pembela kebeneran",
		Rating:      8,
		Image:       "",
	}

	movieRepository.Mock.On("Update", request).Return(expectedResponse)

	app := fiber.New()

	app.Patch("/movies", movieController.UpdateMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPatch, "/movies", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusInternalServerError, httpResponse.Code)

}

func TestUpdateMovie_ValidationFailed(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  500,
		Error: errors.New(constant.ErrTestResponse),
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.UpdateMovieRequest{
		Id:          1,
		Title:       "",
		Description: "Pahlawan super pembela kebeneran",
		Rating:      8,
		Image:       "",
	}

	movieRepository.Mock.On("Update", request).Return(expectedResponse)

	app := fiber.New()

	app.Patch("/movies", movieController.UpdateMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPatch, "/movies", bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

}

func TestUpdateMovie_InvalidJSONParameters(t *testing.T) {

	expectedResponse := common.StatusResponse{
		Code:  500,
		Error: errors.New(constant.ErrTestResponse),
	}

	validate := validator.New()
	movieRepository := &mck.MovieRepositoryMock{Mock: mock.Mock{}}
	movieController := controller.NewMovieController(validate, movieRepository)

	request := common.UpdateMovieRequest{
		Id:          1,
		Title:       "Power Ranger",
		Description: "Pahlawan super pembela kebeneran",
		Rating:      8,
		Image:       "",
	}

	movieRepository.Mock.On("Update", request).Return(expectedResponse)

	app := fiber.New()

	app.Patch("/movies", movieController.UpdateMovie)

	// Create a request
	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPatch, "/movies", bytes.NewBuffer(requestBody))

	// comment the code below to make apicall cant read the json request propertly
	// req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req)

	httpResponse := common.HttpResponse{}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &httpResponse)

	fmt.Println("result :", httpResponse)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

}
