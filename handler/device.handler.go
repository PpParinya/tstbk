package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/database"
	"github.com/tst/backend/model/entity"
	"github.com/tst/backend/model/request"
)

func GetDeviceByUser(ctx *fiber.Ctx) error {

	// var userID = -9223372036854771339 //server
	// var userID = -9223372036854775800 //local

	var userID = ctx.Query("userId")

	var device []entity.Devices
	database.DB.Debug().Table("Devices").Where("userID = ?", userID).Find(&device)
	return ctx.JSON(device)
}

func AddDevice(ctx *fiber.Ctx) error {

	var device = new(request.Devices)
	if err := ctx.BodyParser(device); err != nil {
		return err
	}

	device.AddedOn = time.Now()
	device.ModifiedOn = time.Now()

	errCreateDevice := database.DB.Debug().Table("Devices").Create(&device).Error
	if errCreateDevice != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to insert Device",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    device,
	})

}

func UpdateDevice(ctx *fiber.Ctx) error {

	var deviceId = ctx.Query("deviceId")

	var device = new(request.Devices)
	if err := ctx.BodyParser(device); err != nil {
		return err
	}
	
	device.ModifiedOn = time.Now()

	errCreateDevice := database.DB.Debug().Where("DeviceID = ?", deviceId).Updates(&device).Error
	if errCreateDevice != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to insert Device",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    device,
	})

}

func DeleteDevice(ctx *fiber.Ctx) error {

	var deviceId = ctx.Query("deviceId")
	var device = new(request.Devices)

	errDeleteDevice := database.DB.Debug().Table("Devices").Where("DeviceID = ?", deviceId).Delete(&device).Error
	if errDeleteDevice != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to insert Device",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})

}
