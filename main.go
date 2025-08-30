package main

import (
	"fiber/config"
	"fiber/routes"
	"fiber/migrations"
	"fiber/seeders"

	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	config.ConnectDB()
	migrations.RemoveCityColumn()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	seeders.userSeeder()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
