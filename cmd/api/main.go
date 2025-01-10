package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santicano2/ticket-booking/handlers"
	"github.com/santicano2/ticket-booking/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "TicketBookingAPI",
		ServerHeader: "fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}