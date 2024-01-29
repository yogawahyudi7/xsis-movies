package repository_test

import (
	"fmt"
	"testing"

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
	mock.ExpectQuery("SELECT (.+) FROM \"movies\"").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "desciption", "rating", "image", "column:created_at", "column:updated_at", "column:deleted_at"}).
			AddRow(1, "Power Ranger", "Pahlawan Super", 7, "", "2024-01-28 21:25:23.766 +0700", "2024-01-28 21:25:23.766 +0700", nil))

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
	mock.ExpectQuery("SELECT (.+) FROM \"movies\"").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "desciption", "rating", "image", "column:created_at", "column:updated_at", "column:deleted_at"}))

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
	mock.ExpectQuery("SELECT (.+) FROM \"movie\"").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "desciption", "rating", "image", "column:created_at", "column:updated_at", "column:deleted_at"}))

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetAll()

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	assert.Error(t, response.Error)
}
