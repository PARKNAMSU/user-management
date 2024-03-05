package auth_router

import "github.com/gofiber/fiber/v2"

func UserRouting(v1 fiber.Router) {
	v1.Route("/auth", func(router fiber.Router) {
		router.Post("/issuToken", func(c *fiber.Ctx) error {
			return nil
		})
		router.Post("/variableToken", func(c *fiber.Ctx) error {
			return nil
		})
		router.Post("/getInformationByToken", func(c *fiber.Ctx) error {
			return nil
		})
		router.Post("/extendToken", func(c *fiber.Ctx) error {
			return nil
		})
		router.Post("/deleteToken", func(c *fiber.Ctx) error {
			return nil
		})
	})
}
