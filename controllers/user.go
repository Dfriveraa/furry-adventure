package controllers

import (
	"github.com/Sht97/furry-adventure/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.Status(201).JSON(user)
}
