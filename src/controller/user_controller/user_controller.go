package user_controller

import (
	"github.com/PARKNAMSU/user-management/constant"
	"github.com/PARKNAMSU/user-management/src/controller"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	action constant.UrlAction
}

func GetUserController() controller.Controller {
	return userController{}
}

func (c userController) Controller(fiber *fiber.Ctx) {
	switch c.action {
	// ... todo
	}
}
