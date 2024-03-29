package router

import (
	"movies-xsis/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterMoviePath(app *fiber.App, e *controller.MovieController) {
	app.Get("/movies", e.GetAllMovies)
	app.Post("/movies", e.AddMovie)
	app.Get("/movies/:id", e.GetMovie)
	app.Delete("/movies/:id", e.DeleteMovie)
	app.Patch("/movies/:id", e.UpdateMovie)
}
