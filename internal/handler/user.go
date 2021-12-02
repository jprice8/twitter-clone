package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/model"
	"github.com/jprice8/twitter-clone/internal/shared/database"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Fetch all users
func FetchUsers(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Select all users from database
		rows, err := db.Query("SELECT id, name, email, password, created_at FROM users")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		defer rows.Close()
		result := model.Users{}

		for rows.Next() {
			user := model.User{}
			if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
				return err // Exit if error
			}

			// Append User to Users
			result.Users = append(result.Users, user)
		}

		return c.JSON(result)
	}
}

// GetUser
func GetUser(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user model.User
		row := db.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE id = $1", id)
		if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return err // Exit if error
		}
		return c.JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
	}
}

// Create new user
func CreateUser(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type NewUser struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Missing or invalid fields", "data": err})
		}

		hash, err := hashPassword(user.Password)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
		}

		user.Password = hash
		res, err := db.Query("INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Password, time.Now())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
		}

		log.Println(res)

		newUser := NewUser{
			Name:  user.Name,
			Email: user.Email,
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Successfully created user", "data": newUser})
	}
}

// Delete existing user
func DeleteUser(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := db.Query("DELETE FROM users WHERE id = $1", id)
		if err != nil {
			return c.JSON(fiber.Map{"status": "error", "message": "Could not delete user from database", "data": err})
		}

		log.Println(res)

		return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": id})
	}
}
