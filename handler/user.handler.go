package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/database"
	"github.com/tst/backend/model/entity"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	userInfos := ctx.Locals("userID")
	fmt.Println(userInfos)

	var users []entity.Users

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}
