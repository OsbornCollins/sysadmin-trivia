package handlers

import (
	"github.com/OsbornCollins/sysadmin-trivia/database"
	"github.com/OsbornCollins/sysadmin-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	// Render a template named 'index.html' with content
	return c.Render("index", fiber.Map{
		"Title":       "Hello, World!",
		"Description": "This is a template.",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}
