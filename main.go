package main

import (
	"fmt"
	"github.com/Sht97/furry-adventure/core"
	"github.com/Sht97/furry-adventure/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := core.Configuration

	app := fiber.New()
	routes.SetupRoutes(app)

	listen := ":" + config.Port
	fmt.Println("Server started on port", listen)

	err := app.Listen(listen)
	fmt.Print("Error: ", err)
	if err != nil {
		panic(err)
	}
}
