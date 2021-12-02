package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/handler"
	"github.com/jprice8/twitter-clone/internal/shared/database"
)

func ApiRoutes(r fiber.Router, db database.Database) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})
}

// Auth
func AuthRoutes(r fiber.Router, db database.Database) {
	// Login user
	r.Post("/login", handler.Login(db))
}

// User
func UserRoutes(r fiber.Router, db database.Database) {
	// Get all users
	r.Get("/users", handler.FetchUsers(db))
	// Get user
	r.Get("/:id", handler.GetUser(db))
	// Create new user
	r.Post("/", handler.CreateUser(db))
	// Delete existing user
	r.Delete("/:id", handler.DeleteUser(db))
}

// Product
func ProductRoutes(r fiber.Router, db database.Database) {
	// TODO product routes
	// Get all products
	// Get product
}