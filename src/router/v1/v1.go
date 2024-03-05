package v1

import (
	"github.com/PARKNAMSU/user-management/src/middleware"
	"github.com/PARKNAMSU/user-management/src/router/v1/user_router"
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	app.Route("/api/v1", func(router fiber.Router) {
		router.Use(middleware.ApiValidation)
		user_router.UserRouting(router)
	})
}
