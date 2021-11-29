package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "log"
	"github.com/jprice8/twitter-clone/internal/shared/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Import .env file to environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Print(err)
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%v:%v/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	// Create db
	db, err := database.New(ctx, dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
