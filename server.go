package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi ✌️")
	})
	app.Post("/register", Register)
	app.Post("/login", Login)

	app.Listen(":4000")
}

func Register(c *fiber.Ctx) error {
	return nil
}

func Login(ctx *fiber.Ctx) error {
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var c credentials
	err := ctx.BodyParser(&c)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = c.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("supersecret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{"status": "success", "token": t})
}
