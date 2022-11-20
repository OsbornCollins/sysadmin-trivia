// Filename: cmd/main.go

package main

import (
	"github.com/OsbornCollins/sysadmin-trivia/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html" // add html engine
)

func main() {
	// Establish database connection
	database.ConnectDb()

	// Establish web application connection
	//app := fiber.New()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./views")
	// Function that handles routes
	setupRoutes(app)

	app.Listen(":3000")
}
