package user_router

import "github.com/gofiber/fiber/v2"

func UserRouterV1(router fiber.Router) {
	router.Route("/user", func(userRouter fiber.Router) {
		userRouter.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("user test")
		})
	})
}
