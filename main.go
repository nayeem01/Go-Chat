package main

import (
	"fmt"
	"log"

	"github.com/nayeem01/Go-Chat/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConncetDB()
	app := fiber.New()

	app.Get("/api/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ hi ")
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":8080"))
}
