package routes

import (
	"fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "🚀 Fiber App",
			"Name":  "Anand",
		})
	})

	app.Get("/about",controllers.About)

}
