package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nayeem01/Go-Chat/pkg/controllers"
)

func SetupUserRoutes(app *fiber.App) {

	app.Post("/api/v1/user", controllers.CreateUser)
	app.Get("/api/v1/user", controllers.GetUsers)
	app.Get("/api/v1/user/:id", controllers.GetUser)
	app.Put("/api/v1/user/update/:id", controllers.UpdateUser)

}
