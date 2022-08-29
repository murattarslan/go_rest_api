package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/get-all-base", getAllBase)
	app.Post("/add-base", addBase)
	app.Put("/update-base/:id", updateBase)
	app.Delete("/delete-base", deleteBase)

	app.Use(func(c *fiber.Ctx) error {
		response := fmt.Sprintf("%s%s is not found =(", c.BaseURL(), c.OriginalURL())
		return c.Status(404).SendString(response)
	})
	app.Listen(":8080")
}
