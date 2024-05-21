package routes

import (
	"github.com/ItzRyz/squareville-api/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r *fiber.App) {
	r.Post("/login", controllers.Login)
}
