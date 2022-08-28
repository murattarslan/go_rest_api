package main

import (
	"github.com/gofiber/fiber/v2"
)

func updateBase(c *fiber.Ctx) error {

	id := c.Params("id")
	return c.Status(200).SendString(id)
}
