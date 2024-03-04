package common

import (
	"github.com/gofiber/fiber/v2"
)

func GetReqHeader(ctx *fiber.Ctx, key string) string {
	data := ctx.GetReqHeaders()[key]
	if len(data) <= 0 {
		return ""
	}
	return data[0]
}
