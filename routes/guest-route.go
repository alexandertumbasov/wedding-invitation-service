package routes

import (
	"github.com/gofiber/fiber/v2"
	"wedding-invitation-service/controllers"
)

func GuestRoute(app *fiber.App) {
	app.Post("/guest", controllers.CreateGuest)
}
