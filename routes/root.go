package routes

import (
	"github.com/Sht97/furry-adventure/models"
	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(models.HealthCheck)
}

func setupRootRoute(app fiber.Router) {
	app.Get("/", healthCheck)
}
