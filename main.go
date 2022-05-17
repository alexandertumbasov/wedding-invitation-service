package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"wedding-invitation-service/routes"
)

func main() {
	app := fiber.New()

	routes.GuestRoute(app)

	log.Fatal(app.Listen("127.0.0.1:8080"))
}
