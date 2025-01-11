package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/santicano2/ticket-booking/config"
	"github.com/santicano2/ticket-booking/db"
	"github.com/santicano2/ticket-booking/handlers"
	"github.com/santicano2/ticket-booking/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "TicketBookingAPI",
		ServerHeader: "fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}