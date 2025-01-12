package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/santicano2/ticket-booking/config"
	"github.com/santicano2/ticket-booking/db"
	"github.com/santicano2/ticket-booking/handlers"
	"github.com/santicano2/ticket-booking/middlewares"
	"github.com/santicano2/ticket-booking/repositories"
	"github.com/santicano2/ticket-booking/services"
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
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// Handlers
	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}