package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberUse(app *fiber.App) fiber.Router {

	app.Use(recover.New())

	return app
}
