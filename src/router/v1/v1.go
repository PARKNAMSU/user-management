package v1

import (
	"github.com/PARKNAMSU/user-management/src/middleware"
	"github.com/PARKNAMSU/user-management/src/router/v1/user_v1"
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	app.Route("/api/v1", func(router fiber.Router) {
		router.Use(middleware.ApiValidation)
		user_v1.UserRouting(router)
	})
}
