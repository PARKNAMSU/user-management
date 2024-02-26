package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load()
)

var (
	port = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Listen(port)
}
