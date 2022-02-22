package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"app/api"
)

func main() {

	app := fiber.New()

	// Connect to DB

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	// Set Routes & Middleware
	api.SetupAPI(app)

	// Start API
	fmt.Println("\nStarting app at http://localhost:5000")
	log.Fatal(app.Listen(":5000"))
}
