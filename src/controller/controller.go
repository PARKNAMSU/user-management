package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Controller(c *fiber.Ctx)
}
