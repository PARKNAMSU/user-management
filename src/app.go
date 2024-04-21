package src

import (
	"errors"
	"log"

	"github.com/PARKNAMSU/user-management/configs/app_configs"
	user_router "github.com/PARKNAMSU/user-management/src/user/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	_   = godotenv.Load()
	app *fiber.App
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
		for _, accept := range app_configs.CtxAcceptList {
			ctx.Accepts(accept...)
		}
		return ctx.Next()
	})

	app.Get("/health/check", func(ctx *fiber.Ctx) error {
		return ctx.SendString("alive")
	})

	app.Route("/v1", func(router fiber.Router) {
		user_router.UserRouterV1(router)
	})

	log.Println(app_configs.Port)

	err := app.Listen(app_configs.Port)

	if err != nil {
		app.Listen(":80")
	}
}
