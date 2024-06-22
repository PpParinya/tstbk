package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/database"
	"github.com/tst/backend/model/entity"
	"github.com/tst/backend/model/request"
)

func GetMarkerByUser(ctx *fiber.Ctx) error {

	var userID = ctx.Query("userId")

	var device []entity.CustomMarkers
	database.DB.Debug().Table("CustomMarkers").Where("userID = ?", userID).Find(&device)
	return ctx.JSON(device)
}

func AddMarker(ctx *fiber.Ctx) error {

	var customMarkers = new(request.CustomMarkers)
	if err := ctx.BodyParser(customMarkers); err != nil {
		return err
	}

	userInfos := ctx.Locals("userID").(string)

	customMarkers.CreatedBy = userInfos
	customMarkers.CreatedOn = time.Now()

	errCreateDevice := database.DB.Debug().Table("customMarkers").Create(&customMarkers).Error
	if errCreateDevice != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to insert customMarkers",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    customMarkers,
	})

}

func UpdateMarker(ctx *fiber.Ctx) error {

	return nil

}

func DeleteMarker(ctx *fiber.Ctx) error {

	var markerId = ctx.Query("MarkerId")
	var marker = new(request.CustomMarkers)

	errDeleteDevice := database.DB.Debug().Table("customMarkers").Where("MarkerID = ?", markerId).Delete(&marker).Error
	if errDeleteDevice != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to Delete Marker",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})

}
