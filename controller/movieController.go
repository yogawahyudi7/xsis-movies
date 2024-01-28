package controller

import (
	"movies-xsis/common"
	"movies-xsis/constant"
	"movies-xsis/repository"

	"github.com/gofiber/fiber/v2"
)

type MovieController struct {
	Movie repository.MovieRepositoryInterface
}

func NewMovieController(movie repository.MovieRepositoryInterface) *MovieController {
	return &MovieController{
		Movie: movie,
	}
}

func (get *MovieController) GetAllMovie(ctx *fiber.Ctx) error {

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
