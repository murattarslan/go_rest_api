package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func addBase(c *fiber.Ctx) error {

	body := new(Body)

	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString(string(err.Error()))
	}
	response := Response{Result{body.Name, body.Id}, ResultMessage{"body has been created", "success", "200"}}
	jsonString, err := json.Marshal(response)
	if err != nil {
		return c.Status(500).SendString(string(err.Error()))
	}

	return c.Status(200).SendString(string(jsonString))
}
