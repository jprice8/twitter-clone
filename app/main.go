package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "log"
	"github.com/jprice8/twitter-clone/internal/server"
	"github.com/jprice8/twitter-clone/internal/shared/database"
	"github.com/jprice8/twitter-clone/internal/shared/webserver"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// // User struct
// type User struct {
// 	ID        int64    `json:"id"`
// 	Name      string    `json:"name"`
// 	Email     string    `json:"email"`
// 	Password  string    `json:"password"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// // Users struct
// type Users struct {
// 	Users []User `json:"users"`
// }

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

	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello world!")
	// })

	// // Get all users from PGSQL
	// app.Get("/users", func(c *fiber.Ctx) error {
	// 	// Select all users from database
	// 	rows, err := db.Query("SELECT id, name, email, password, created_at FROM users")
	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}
	// 	defer rows.Close()
	// 	result := Users{}

	// 	for rows.Next() {
	// 		user := User{}
	// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
	// 			return err // Exit if error
	// 		}

	// 		// Append User to Users
	// 		result.Users = append(result.Users, user)
	// 	}
	// 	// Return Users in JSON format
	// 	return c.JSON(result)
	// })

	// // Add user to PGSQL
	// app.Post("/users", func(c *fiber.Ctx) error {
	// 	// New User struct
	// 	u := new(User)

	// 	// Parse body into struct
	// 	if err := c.BodyParser(u); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Insert User into database
	// 	res, err := db.Query("INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4)", u.Name, u.Email, u.Password, time.Now())
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Print result
	// 	log.Println(res)

	// 	// Return User in JSON
	// 	return c.Status(201).JSON(u)
	// })

	// router.SetupRoutes(app, db)

	// app.Listen(":8080")
	webserver := webserver.New(fiber.Config{})

	server := server.New(webserver, db)

	server.Listen()
}
