package main

import (
	"log"

	"github.com/nayeem01/Go-Chat/pkg/database"
	"github.com/nayeem01/Go-Chat/pkg/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConncetDB()
	app := fiber.New()

	routers.SetupUserRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
