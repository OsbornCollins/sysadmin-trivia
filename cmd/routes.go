// Filename: cmd/routes.go

package main

import (
	"github.com/OsbornCollins/sysadmin-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
	app.Get("/fact", handlers.ListFacts)
}
