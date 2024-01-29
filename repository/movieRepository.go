package repository

import (
	"errors"
	"movies-xsis/common"
	"movies-xsis/constant"
	"movies-xsis/model"

	"gorm.io/gorm"
)

type MovieRepositoryInterface interface {
	GetAll() ([]model.Movie, common.StatusResponse)
	Add(parameter model.Movie) (response common.StatusResponse)
	GetById(parameter int) (model.Movie, common.StatusResponse)
	Delete(parameter int) (response common.StatusResponse)
	Update(parameter model.Movie) (response common.StatusResponse)
}

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (movie *MovieRepository) GetAll() ([]model.Movie, common.StatusResponse) {

	movies := []model.Movie{}
	response := common.StatusResponse{}

	query := movie.db.Debug()
	query = query.Find(&movies)

	if query.Error != nil {

		response.Code = 500
		response.Error = query.Error

		return movies, response
	}

	if len(movies) < 1 {
		response.Code = 404
		response.Error = errors.New(constant.ErrDataEmpty)

		return movies, response
	}

	return movies, response
}

func (movie *MovieRepository) Add(parameter model.Movie) (response common.StatusResponse) {

	query := movie.db.Debug()
	query = query.Create(&parameter)

	if query.Error != nil {

		response.Code = 500
		response.Error = query.Error

		return response
	}

	if query.RowsAffected < 1 {
		response.Code = 500
		response.Error = errors.New(constant.ErrRowAffected)

		return response
	}

	return response
}

func (movie *MovieRepository) GetById(parameter int) (model.Movie, common.StatusResponse) {

	movies := model.Movie{}
	response := common.StatusResponse{}

	query := movie.db.Debug()
	query = query.Where("id = ?", parameter)
	query = query.Find(&movies)

	if query.Error != nil {

		response.Code = 500
		response.Error = query.Error

		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			response.Code = 404
		}

		return movies, response
	}

	if movies.Id == 0 {
		response.Code = 404
		response.Error = errors.New(constant.ErrDataEmpty)

		return movies, response
	}

	return movies, response
}

func (movie *MovieRepository) Delete(parameter int) (response common.StatusResponse) {

	deleteMovie := model.Movie{}

	query := movie.db.Debug()
	query = query.Where("id = ?", parameter)
	query = query.Delete(&deleteMovie)

	if query.Error != nil {

		response.Code = 500
		response.Error = query.Error

		return response
	}

	if query.RowsAffected < 1 {
		response.Code = 500
		response.Error = errors.New(constant.ErrRowAffected)

		return response
	}

	return response
}

func (movie *MovieRepository) Update(parameter model.Movie) (response common.StatusResponse) {

	query := movie.db.Debug()
	query = query.Where("id = ?", parameter.Id)
	query = query.Updates(&parameter)

	if query.Error != nil {

		response.Code = 500
		response.Error = query.Error

		return response
	}

	if query.RowsAffected < 1 {
		response.Code = 500
		response.Error = errors.New(constant.ErrRowAffected)

		return response
	}

	return response
}
