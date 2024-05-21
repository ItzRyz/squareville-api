package main

import (
	"fmt"

	"github.com/ItzRyz/squareville-api/src/database"
	"github.com/ItzRyz/squareville-api/src/database/migrations"
	"github.com/ItzRyz/squareville-api/src/lib"
	"github.com/ItzRyz/squareville-api/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.DatabaseInit()
	migrations.Migration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello",
		})
	})

	routes.RouteInit(app)
	port := lib.GetEnvVar("BASE_PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}
