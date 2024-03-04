package middleware

import (
	"errors"
	"os"

	"github.com/PARKNAMSU/user-management/src/tool/common"
	"github.com/gofiber/fiber/v2"
)

func ApiValidation(ctx *fiber.Ctx) error {
	if os.Getenv("API_KEY") != common.GetReqHeader(ctx, "X-Api-Key") {
		return errors.New("not allowed")
	}
	return ctx.Next()
}
