package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/model"
	"github.com/jprice8/twitter-clone/internal/shared/config"
	"github.com/jprice8/twitter-clone/internal/shared/database"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string, db database.Database) (*model.User, error) {
	var user model.User
	row := db.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE email = $1", e)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return nil, err // Exit if error
	}
	return &user, nil
}

func Login(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type LoginInput struct {
			Identity string `json:"identity"`
			Password string `json:"password"`
		}

		var input LoginInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
		}
		identity := input.Identity
		pass := input.Password

		user, err := getUserByEmail(identity, db)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
		}

		if !CheckPasswordHash(pass, user.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name
		claims["user_id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(config.Config("SECRET")))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Login Successful", "data": t})
	}
}
