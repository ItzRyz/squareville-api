package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	UserRoutes(r)
	AuthRoutes(r)
}
