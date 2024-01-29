package repository_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"movies-xsis/constant"
	"movies-xsis/model"
	"movies-xsis/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
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

func TestGetAllMovies_QueryError(t *testing.T) {

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

func TestGetMovie_DataFound(t *testing.T) {

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
	mock.ExpectQuery("SELECT (.+) FROM \"movies\" WHERE id = ?").WillReturnRows(rows).
		WithArgs(1)

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetById(1)

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	assert.NoError(t, response.Error)
}

func TestGetMovie_RecordNotFound(t *testing.T) {

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

	mock.ExpectQuery("SELECT (.+) FROM \"movies\" WHERE id = ?").WillReturnError(gorm.ErrRecordNotFound)

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetById(1)

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	// assert.NoError(t, response.Error)
	assert.Error(t, response.Error)
}

func TestGetMovie_DataEmpty(t *testing.T) {

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
	mock.ExpectQuery("SELECT (.+) FROM \"movies\" WHERE id = ?").WillReturnRows(rows).
		WithArgs(1)

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetById(1)

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	assert.Error(t, response.Error)
}

func TestGetMovie_QueryError(t *testing.T) {

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
	mock.ExpectQuery("SELECT (.+) FROM \"movie\" WHERE id = ?").WillReturnRows(rows).
		WithArgs(1)

	movieRepository := repository.NewMovieRepository(gormDB)

	data, response := movieRepository.GetById(1)

	fmt.Println("Response : ", response)
	fmt.Println("data : ", data)

	assert.Error(t, response.Error)
}

// IMPLEMENT MYSQL DRIVER FOR SUPPORT ADD REPOSITORY TESTING
func TestAddMovie_DataSaved(t *testing.T) {

	// Setup GORM with sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	movieRepository := repository.NewMovieRepository(gormDB)

	time := time.Now()

	// Data movie yang akan diuji
	query := model.Movie{
		Id:          0,
		Title:       "Movie Title",
		Description: "Movie Description",
		Rating:      8.5,
		Image:       "movie.jpg",
		CreatedAt:   &time,
		UpdatedAt:   &time,
		DeletedAt:   nil,
	}

	// Expect Exec to be called with a raw query
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `movies`").
		WithArgs(
			query.Title,
			query.Description,
			query.Rating,
			query.Image,
			query.CreatedAt,
			query.UpdatedAt,
			query.DeletedAt,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	response := movieRepository.Add(query)

	assert.NoError(t, response.Error)
}

func TestAddMovie_QueryError(t *testing.T) {

	// Setup GORM with sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	movieRepository := repository.NewMovieRepository(gormDB)

	time := time.Now()

	// Data movie yang akan diuji
	query := model.Movie{
		Id:          0,
		Title:       "Movie Title",
		Description: "Movie Description",
		Rating:      8.5,
		Image:       "movie.jpg",
		CreatedAt:   &time,
		UpdatedAt:   &time,
		DeletedAt:   nil,
	}

	// Expect Exec to be called with a raw query
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `moviess`").
		WithArgs(
			query.Title,
			query.Description,
			query.Rating,
			query.Image,
			query.CreatedAt,
			query.UpdatedAt,
			query.DeletedAt,
		).WillReturnError(errors.New(constant.ErrTestResponse))
	mock.ExpectCommit()

	response := movieRepository.Add(query)

	assert.Error(t, response.Error)
}

func TestAddMovie_RowAffectedIsNull(t *testing.T) {

	// Setup GORM with sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	movieRepository := repository.NewMovieRepository(gormDB)

	time := time.Now()

	// Data movie yang akan diuji
	query := model.Movie{
		Id:          0,
		Title:       "Movie Title",
		Description: "Movie Description",
		Rating:      8.5,
		Image:       "movie.jpg",
		CreatedAt:   &time,
		UpdatedAt:   &time,
		DeletedAt:   nil,
	}

	// Expect Exec to be called with a raw query
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `movies`").
		WithArgs(
			query.Title,
			query.Description,
			query.Rating,
			query.Image,
			query.CreatedAt,
			query.UpdatedAt,
			query.DeletedAt,
		).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectCommit()

	response := movieRepository.Add(query)

	assert.Error(t, response.Error)
}

// IMPLEMENT MYSQL DRIVER FOR SUPPORT UPDATE REPOSITORY TESTING
func TestUpdateMovie_DataSaved(t *testing.T) {

	// Setup GORM with sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	movieRepository := repository.NewMovieRepository(gormDB)

	time := time.Now()

	// Data movie yang akan diuji
	query := model.Movie{
		Id:          1,
		Title:       "Movie Title",
		Description: "Movie Description",
		Rating:      8.5,
		Image:       "movie.jpg",
		CreatedAt:   &time,
		UpdatedAt:   &time,
		DeletedAt:   nil,
	}

	// Expect Exec to be called with a raw query
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `movies` SET").
		WithArgs(
			query.Title,
			query.Description,
			query.Rating,
			query.Image,
			query.CreatedAt,
			query.UpdatedAt,
			// query.DeletedAt,
			query.Id,
			query.Id, // NEED THIS! THIS IS ANOMALI IN SQLMOCK
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	response := movieRepository.Update(query)

	assert.NoError(t, response.Error)
}

func TestUpdateMovie_QueryError(t *testing.T) {

	// Setup GORM with sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	movieRepository := repository.NewMovieRepository(gormDB)

	time := time.Now()

	// Data movie yang akan diuji
	query := model.Movie{
		Id:          1,
		Title:       "Movie Title",
		Description: "Movie Description",
		Rating:      8.5,
		Image:       "movie.jpg",
		CreatedAt:   &time,
		UpdatedAt:   &time,
		DeletedAt:   nil,
	}

	// Expect Exec to be called with a raw query
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `movies` SET").
		WithArgs(
			query.Title,
			query.Description,
			query.Rating,
			query.Image,
			query.CreatedAt,
			query.UpdatedAt,
			// query.DeletedAt,
			query.Id,
			query.Id, // NEED THIS! THIS IS ANOMALI IN SQLMOCK
		).WillReturnError(errors.New(constant.ErrTestResponse))
	mock.ExpectCommit()

	response := movieRepository.Update(query)

	assert.Error(t, response.Error)
}

func TestUpdateMovie_RowAffectedIsNull(t *testing.T) {

	// Setup GORM with sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	movieRepository := repository.NewMovieRepository(gormDB)

	time := time.Now()

	// Data movie yang akan diuji
	query := model.Movie{
		Id:          1,
		Title:       "Movie Title",
		Description: "Movie Description",
		Rating:      8.5,
		Image:       "movie.jpg",
		CreatedAt:   &time,
		UpdatedAt:   &time,
		DeletedAt:   nil,
	}

	// Expect Exec to be called with a raw query
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `movies` SET").
		WithArgs(
			query.Title,
			query.Description,
			query.Rating,
			query.Image,
			query.CreatedAt,
			query.UpdatedAt,
			// query.DeletedAt,
			query.Id,
			query.Id, // NEED THIS! THIS IS ANOMALI IN SQLMOCK
		).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectCommit()

	response := movieRepository.Update(query)

	assert.Error(t, response.Error)
}
