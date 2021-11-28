package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/shared/config"
)

func main() {
	cfgPath, err := filepath.Abs("./configs/.env")
	if err != nil {
		log.Fatal(err)
	}

	config.Load(cfgPath)
	fmt.Springf("cfg Path: %c", cfgPath)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
