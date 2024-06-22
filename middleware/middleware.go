package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/utils"
)

func Auth(ctx *fiber.Ctx) error {

	token := ctx.Get("tstToken")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// role := claims["role"].(string)
	// if role != "admin" {
	// 	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
	// 		"message": "forbidde access",
	// 	})
	// }

	ctx.Locals("userInfo", claims)
	ctx.Locals("userID", claims["userID"])

	return ctx.Next()
}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
