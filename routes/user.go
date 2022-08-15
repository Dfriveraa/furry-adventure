package routes

import (
	"github.com/Sht97/furry-adventure/controllers"
	"github.com/gofiber/fiber/v2"
)

func setupUsersRoutes(app fiber.Router) {
	app.Post("/user", controllers.CreateUser)	
}