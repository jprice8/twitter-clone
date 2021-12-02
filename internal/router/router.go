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
	// Get all products
	r.Get("/", handler.GetAllProducts(db))
	// Get product by categories
	r.Get("/category/:categoryId", handler.GetProductsByCategory(db))
	// Lookahead for autocomplete search
	r.Get("/lookAhead/:prefix", handler.LookAhead(db))

	// Create new product
	r.Post("/", handler.CreateProduct(db))
	// Categorize a product
	r.Post("/categorize/:productId/:categoryId", handler.CategorizeProduct(db))
}

// Category
func CategoryRoutes(r fiber.Router, db database.Database) {
	// Create new category
	r.Post("/", handler.CreateCategory(db))
}