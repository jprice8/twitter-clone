package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "log"
	// "path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	// "github.com/jprice8/twitter-clone/internal/shared/config"
)

func main() {
	// cfgPath, err := filepath.Abs("./configs/.env")
	// if err != nil {
	// log.Fatal(err)
	// }

	// config.Load(cfgPath)
	// fmt.Printf("cfg Path: %v", cfgPath)

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

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Failed to init db.", err)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
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
