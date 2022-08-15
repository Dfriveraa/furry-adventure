package routes

import (
	"github.com/Sht97/furry-adventure/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupFilesRoutes(app fiber.Router) {
	app.Get("/files/save", controllers.SaveFile)
	app.Get("/files", controllers.GetFiles)

}
