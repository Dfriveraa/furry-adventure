package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx) error {
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

func GetFiles(c *fiber.Ctx) error {
	return c.Status(200).SendString("Files")
}
