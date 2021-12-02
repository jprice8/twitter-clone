package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/model"
	"github.com/jprice8/twitter-clone/internal/shared/database"
)

// Create new category
func CreateCategory(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		category := new(model.Category)
		if err := c.BodyParser(category); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse category", "data": err})
		}

		res, err := db.Query("INSERT INTO categories (name, description) VALUES ($1, $2)", category.Name, category.Description)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create new category in db", "data": err})
		}

		log.Println(res)

		return c.JSON(fiber.Map{"status": "success", "message": "Successfully created category", "data": category.Name})
	}
}