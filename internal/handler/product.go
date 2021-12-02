package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/model"
	"github.com/jprice8/twitter-clone/internal/shared/database"
)

type ProductByCategory struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	CategoryNames string  `json:"category_names"`
}

type ProductsByCategory struct {
	ProductsByCategory []ProductByCategory `json:"productsByCategory`
}

// Get all products
func GetAllProducts(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Select all products
		rows, err := db.Query("SELECT * FROM products")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		defer rows.Close()
		result := model.Products{}

		for rows.Next() {
			product := model.Product{}
			if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
				return err // Exit if error
			}

			// Append product to Products
			result.Products = append(result.Products, product)
		}

		return c.JSON(result)
	}
}

// Get products by category
func GetProductsByCategory(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {

		categoryId := c.Params("categoryId")
		// Select products by category id
		rows, err := db.Query(`
		SELECT products.id,
				products.name,
				products.price,
				string_agg(DISTINCT categories.name, ',') AS category_names
		FROM products
		JOIN product_categories ON products.id = product_categories.product_id
		JOIN categories ON categories.id = product_categories.category_id
		WHERE categories.id = $1
		GROUP BY products.id
		`, categoryId)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Query failed on productsByCategory", "data": err})
		}

		defer rows.Close()
		result := ProductsByCategory{}

		for rows.Next() {
			productByCategory := ProductByCategory{}
			if err := rows.Scan(&productByCategory.ID, &productByCategory.Name, &productByCategory.Price, &productByCategory.CategoryNames); err != nil {
				return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error while scanning productByCategory rows", "data": err})
			}

			// Append to Products
			result.ProductsByCategory = append(result.ProductsByCategory, productByCategory)
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Successfully returned products by category", "data": result})
	}
}

// Lookahead endpoint for autocomplete
func LookAhead(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		prefix := c.Params("prefix")
		rows, err := db.Query(`
		SELECT * FROM products
		WHERE products.name LIKE $1
		LIMIT 5
		`, prefix)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Query failed on lookAhead", "data": err})
		}

		defer rows.Close()
		result := model.Products{}

		for rows.Next() {
			product := model.Product{}
			if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
				return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error while scanning lookAhead rows", "data": err})
			}

			// Append product to Products
			result.Products = append(result.Products, product)
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Successfully returned products for lookAhead", "data": result})
	}
}

// Create new product
func CreateProduct(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		product := new(model.Product)
		if err := c.BodyParser(product); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse product input", "data": err})
		}

		res, err := db.Query("INSERT INTO products (name, description, price, created_at) VALUES ($1, $2, $3, $4)", product.Name, product.Description, product.Price, time.Now())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create new product in database", "data": err})
		}

		log.Println(res)

		return c.JSON(fiber.Map{"status": "success", "message": "Successfully created new product", "data": product.Name})
	}
}

// Categorize product
func CategorizeProduct(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {

		productId := c.Params("productId")
		categoryId := c.Params("categoryId")
		res, err := db.Query("INSERT INTO product_categories (product_id, category_id) VALUES ($1, $2)", productId, categoryId)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create new product_category in database", "data": err})
		}

		log.Println(res)

		return c.JSON(fiber.Map{"status": "success", "message": "Successfully created new product_category", "data": productId})
	}
}
