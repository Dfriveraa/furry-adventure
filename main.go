package main

import (
	"fmt"

	"github.com/Sht97/furry-adventure/core"
	"github.com/Sht97/furry-adventure/routes"
	"github.com/gofiber/fiber/v2"
)


func saveFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	fmt.Println(file.Filename)
	err = c.SaveFile(file, file.Filename)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.Status(201).SendString("File saved")
}

func main() {

	config := core.Configuration

	app := fiber.New()
	routes.SetupRoutes(app)

	listen := ":" + config.Port
	fmt.Println("Server started on port", listen)

	app.Listen(listen)
}
