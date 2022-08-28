package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func getAllBase(c *fiber.Ctx) error {

	response := Response{Result{"go router", 123}, ResultMessage{"response has been sent", "success", "200"}}
	jsonString, err := json.Marshal(response)
	if err != nil {
		return c.Status(500).SendString(string(err.Error()))
	}

	return c.Status(200).SendString(string(jsonString))
}
