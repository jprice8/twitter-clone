package handler

import (
	// "strconv"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/model"
	"github.com/jprice8/twitter-clone/internal/shared/database"
	// "golang.org/x/crypto/bcrypt"
)

// func hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func validToken(t *jwt.Token, id string) bool {
// 	n, err := strconv.Atoi(id)
// 	if err != nil {
// 		return false
// 	}

// 	claims := t.Claims.(jwt.MapClaims)
// 	uid := int(claims["user_id"].(float64))

// 	return uid == n
// }

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
