package user_router

import "github.com/gofiber/fiber/v2"

func UserRouting(v1 fiber.Router) {
	v1.Route("/user", func(user fiber.Router) {
		user.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("user test")
		})
	})
}
