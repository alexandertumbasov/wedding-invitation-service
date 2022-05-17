package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"wedding-invitation-service/configs"
	"wedding-invitation-service/routes"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.GuestRoute(app)

	log.Fatal(app.Listen(":6060"))
}
