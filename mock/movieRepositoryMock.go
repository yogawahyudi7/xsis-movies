package movie

import (
	"movies-xsis/common"
	"movies-xsis/model"

	"github.com/stretchr/testify/mock"
)

type MovieRepositoryMock struct {
	Mock mock.Mock
}

func (movie *MovieRepositoryMock) GetAll() ([]model.Movie, common.StatusResponse) {
	args := movie.Mock.Called()

	return args.Get(0).([]model.Movie), args.Get(1).(common.StatusResponse)
}

func (movie *MovieRepositoryMock) Add(parameter common.AddMovieRequest) (response common.StatusResponse) {
	args := movie.Mock.Called(parameter)

	return args.Get(0).(common.StatusResponse)
}

func (movie *MovieRepositoryMock) GetById(parameter int) (model.Movie, common.StatusResponse) {
	args := movie.Mock.Called(parameter)

	return args.Get(0).(model.Movie), args.Get(1).(common.StatusResponse)
}

func (movie *MovieRepositoryMock) Delete(parameter int) (response common.StatusResponse) {
	args := movie.Mock.Called(parameter)

	return args.Get(0).(common.StatusResponse)
}

func (movie *MovieRepositoryMock) Update(parameter common.UpdateMovieRequest) (response common.StatusResponse) {
	args := movie.Mock.Called(parameter)

	return args.Get(0).(common.StatusResponse)
}
