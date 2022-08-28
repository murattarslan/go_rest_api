package main

import (
	"github.com/gofiber/fiber/v2"
)

func deleteBase(c *fiber.Ctx) error {

	id := c.Query("id")
	return c.Status(200).SendString(id)
}
