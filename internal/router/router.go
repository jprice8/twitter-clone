package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/handler"
	"github.com/jprice8/twitter-clone/internal/shared/database"
)

// Setup router api
// func SetupRoutes(app *fiber.App, db database.Database) {
// 	//Middleware
// 	api := app.Group("/api", logger.New())
// 	api.Get("/", handler.Hello)
// }

// Define route group
func ApiRoutes(r fiber.Router, db database.Database) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})
}

func AuthRoutes(r fiber.Router, db database.Database) {
	// Login user
	r.Post("/login", handler.Login(db))
}

// Handle user routes
func UserRoutes(r fiber.Router, db database.Database) {
	// Get all users
	r.Get("/users", handler.FetchUsers(db))

	// Get user
	r.Get("/:id", handler.GetUser(db))
}