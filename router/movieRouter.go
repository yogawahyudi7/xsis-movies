package router

import (
	"movies-xsis/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterMoviePath(app *fiber.App, e *controller.MovieController) {
	app.Get("/movies", e.GetAllMovie)
}
