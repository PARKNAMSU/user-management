package src

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/PARKNAMSU/user-management/src/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	_    = godotenv.Load()
	app  *fiber.App
	port = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

func AppInit() {
	app = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			log.Println("err: ", err.Error())
			// Send custom error page
			err = ctx.Status(code).SendString(err.Error())
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			return nil
		},
	})

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Accepts("html")                           // "html"
		ctx.Accepts("text/html")                      // "text/html"
		ctx.Accepts("json", "text")                   // "json"
		ctx.Accepts("application/json")               // "application/json"
		ctx.Accepts("text/plain", "application/json") // "application/json", due to quality
		ctx.Accepts("image/png")                      // ""
		ctx.Accepts("png")                            // "", due to */* without q factor 0 is Not Acceptable
		return nil
	})

	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("alive")
	})

	app.Use(middleware.ApiValidation)

	log.Println(port)

	err := app.Listen(port)

	if err != nil {
		app.Listen(":80")
	}
}
