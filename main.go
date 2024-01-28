package main

import (
	"movies-xsis/config"
	"movies-xsis/controller"
	"movies-xsis/repository"
	"movies-xsis/router"
	"movies-xsis/util"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func main() {

	setup := config.Get()
	postgresql := util.DBConnection(setup)
	movieRepository := repository.NewMovieRepository(postgresql)
	movieController := controller.NewMovieController(movieRepository)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	config.FiberUse(app)

	router.RegisterMoviePath(app, movieController)

	app.Listen(":8080")
}
