package repository_test

import (
	"fmt"
	"testing"
	"time"

	"movies-xsis/model"
	"movies-xsis/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetAllMovies_DataFound(t *testing.T) {

	// Inisialisasi mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database: %v", err)
	}
	defer db.Close()

	// Inisialisasi GORM dengan mock database
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize GORM: %v", err)
	}

	// Pengaturan ekspetasi query dan hasilnya
	time := time.Now()
	moviesData := []model.Movie{
		{
			Id:          1,
			Title:       "Power Ranger",
			Description: "Pahlawan Super",
			Rating:      7,
			Image:       "",
			CreatedAt:   &time,
			UpdatedAt:   &time,
			DeletedAt:   nil,
		},
	}

	// Pengaturan ekspetasi query dan hasilnya
	rows := sqlmock.NewRows([]string{"id", "title", "desciption", "rating", "image", "column:created_at", "column:updated_at", "column:deleted_at"})

	for _, movie := range moviesData {
		rows = rows.AddRow(
			movie.Id, movie.Title, movie.Description, movie.Rating, movie.Image, movie.CreatedAt, movie.UpdatedAt, movie.DeletedAt)
	}
	mock.ExpectQuery("SELECT (.+) FROM \"movies\"").WillReturnRows(rows)

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetAll()

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	assert.NoError(t, response.Error)
}

func TestGetAllMovies_DataEmpty(t *testing.T) {

	// Inisialisasi mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database: %v", err)
	}
	defer db.Close()

	// Inisialisasi GORM dengan mock database
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize GORM: %v", err)
	}

	// Pengaturan ekspetasi query dan hasilnya
	moviesData := []model.Movie{}

	// Pengaturan ekspetasi query dan hasilnya
	rows := sqlmock.NewRows([]string{"id", "title", "desciption", "rating", "image", "column:created_at", "column:updated_at", "column:deleted_at"})

	for _, movie := range moviesData {
		rows = rows.AddRow(
			movie.Id, movie.Title, movie.Description, movie.Rating, movie.Image, movie.CreatedAt, movie.UpdatedAt, movie.DeletedAt)
	}
	mock.ExpectQuery("SELECT (.+) FROM \"movies\"").WillReturnRows(rows)

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetAll()

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	assert.Error(t, response.Error)
}

func TestGetAllMovies_QueryNotMatch(t *testing.T) {

	// Inisialisasi mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database: %v", err)
	}
	defer db.Close()

	// Inisialisasi GORM dengan mock database
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize GORM: %v", err)
	}
	// Pengaturan ekspetasi query dan hasilnya
	time := time.Now()
	moviesData := []model.Movie{
		{
			Id:          1,
			Title:       "Power Ranger",
			Description: "Pahlawan Super",
			Rating:      7,
			Image:       "",
			CreatedAt:   &time,
			UpdatedAt:   &time,
			DeletedAt:   nil,
		},
	}

	// Pengaturan ekspetasi query dan hasilnya
	rows := sqlmock.NewRows([]string{"id", "title", "desciption", "rating", "image", "column:created_at", "column:updated_at", "column:deleted_at"})

	for _, movie := range moviesData {
		rows = rows.AddRow(
			movie.Id, movie.Title, movie.Description, movie.Rating, movie.Image, movie.CreatedAt, movie.UpdatedAt, movie.DeletedAt)
	}
	mock.ExpectQuery("SELECT (.+) FROM \"movie\"").WillReturnRows(rows)

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetAll()

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	assert.Error(t, response.Error)
}
