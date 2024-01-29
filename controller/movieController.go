package controller

import (
	"movies-xsis/common"
	"movies-xsis/constant"
	"movies-xsis/model"
	"movies-xsis/repository"
	"movies-xsis/validator"
	"reflect"

	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type MovieController struct {
	Validate *validate.Validate
	Movie    repository.MovieRepositoryInterface
}

func NewMovieController(validate *validate.Validate, movie repository.MovieRepositoryInterface) *MovieController {
	return &MovieController{
		Validate: validate,
		Movie:    movie,
	}
}

func (get *MovieController) GetAllMovies(ctx *fiber.Ctx) error {

	movieData, movieResponse := get.Movie.GetAll()

	httpResponse := common.HttpResponse{}
	if movieResponse.Error != nil {
		httpResponse.Code = 500
		httpResponse.Message = constant.ServerUnderMaintenance
		httpResponse.Data = nil

		if movieResponse.Code == 404 {
			httpResponse.Code = 404
			httpResponse.Message = constant.DataIsNotAvailabe
		}

		return ctx.JSON(httpResponse)
	}

	data := []common.GetMovieResponse{}

	for _, v := range movieData {
		vData := common.GetMovieResponse{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Image:       v.Image,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}

		data = append(data, vData)
	}

	httpResponse.Code = 200
	httpResponse.Message = constant.DataFound
	httpResponse.Data = data

	return ctx.JSON(httpResponse)
}

func (get *MovieController) AddMovie(ctx *fiber.Ctx) error {

	movie := common.AddMovieRequest{}
	httpResponse := common.HttpResponse{}
	httpResponse.Data = nil

	if err := ctx.BodyParser(&movie); err != nil {
		httpResponse.Code = 400
		httpResponse.Message = constant.InvalidJsonParameters
		return ctx.JSON(httpResponse)
	}

	if err := get.Validate.Struct(movie); err != nil {

		for _, err := range err.(validate.ValidationErrors) {
			field, _ := reflect.TypeOf(movie).FieldByName(err.StructField())

			message := validator.TemplateMessage(field, err)
			httpResponse.Code = 400
			httpResponse.Message = message
			return ctx.JSON(httpResponse)
		}
	}

	req := model.Movie{
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
	}

	movieResponse := get.Movie.Add(req)

	if movieResponse.Error != nil {
		httpResponse.Code = 500
		httpResponse.Message = constant.ServerUnderMaintenance

		return ctx.JSON(httpResponse)
	}

	httpResponse.Code = 200
	httpResponse.Message = constant.DataSaved

	return ctx.JSON(httpResponse)
}

func (get *MovieController) GetMovie(ctx *fiber.Ctx) error {

	httpResponse := common.HttpResponse{}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		httpResponse.Code = 400
		httpResponse.Message = constant.InvalidRouteParameters
		return ctx.JSON(httpResponse)
	}

	movieData, movieResponse := get.Movie.GetById(id)

	if movieResponse.Error != nil {
		httpResponse.Code = 500
		httpResponse.Message = constant.ServerUnderMaintenance
		httpResponse.Data = nil

		if movieResponse.Code == 404 {
			httpResponse.Code = 404
			httpResponse.Message = constant.DataIsNotAvailabe
		}

		return ctx.JSON(httpResponse)
	}

	data := common.GetMovieResponse{
		Id:          movieData.Id,
		Title:       movieData.Title,
		Description: movieData.Description,
		Image:       movieData.Image,
		CreatedAt:   movieData.CreatedAt,
		UpdatedAt:   movieData.UpdatedAt,
	}

	httpResponse.Code = 200
	httpResponse.Message = constant.DataFound
	httpResponse.Data = data

	return ctx.JSON(httpResponse)
}

func (get *MovieController) DeleteMovie(ctx *fiber.Ctx) error {

	httpResponse := common.HttpResponse{}
	httpResponse.Data = nil

	id, err := ctx.ParamsInt("id")
	if err != nil {
		httpResponse.Code = 400
		httpResponse.Message = constant.InvalidRouteParameters
		return ctx.JSON(httpResponse)
	}

	movieResponse := get.Movie.Delete(id)

	if movieResponse.Error != nil {
		httpResponse.Code = 500
		httpResponse.Message = constant.ServerUnderMaintenance

		return ctx.JSON(httpResponse)
	}

	httpResponse.Code = 200
	httpResponse.Message = constant.DataDeleted

	return ctx.JSON(httpResponse)
}

func (get *MovieController) UpdateMovie(ctx *fiber.Ctx) error {

	movie := common.UpdateMovieRequest{}
	httpResponse := common.HttpResponse{}
	httpResponse.Data = nil

	if err := ctx.BodyParser(&movie); err != nil {
		httpResponse.Code = 400
		httpResponse.Message = constant.InvalidJsonParameters
		return ctx.JSON(httpResponse)
	}

	if err := get.Validate.Struct(movie); err != nil {

		for _, err := range err.(validate.ValidationErrors) {
			field, _ := reflect.TypeOf(movie).FieldByName(err.StructField())

			message := validator.TemplateMessage(field, err)
			httpResponse.Code = 400
			httpResponse.Message = message
			return ctx.JSON(httpResponse)
		}
	}

	req := model.Movie{
		Id:          uint(movie.Id),
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
	}
	movieResponse := get.Movie.Update(req)

	if movieResponse.Error != nil {
		httpResponse.Code = 500
		httpResponse.Message = constant.ServerUnderMaintenance

		return ctx.JSON(httpResponse)
	}

	httpResponse.Code = 200
	httpResponse.Message = constant.DataUpdated

	return ctx.JSON(httpResponse)
}
