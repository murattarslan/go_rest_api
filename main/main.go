package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/get-all-base", getAllBase)
	app.Post("/add-base", addBase)
	app.Put("/update-base/:id", updateBase)
	app.Delete("/delete-base", deleteBase)

	app.Listen(":8080")
}
