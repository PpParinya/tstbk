package handler

import (

	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/database"
	"github.com/tst/backend/enum"
	"github.com/tst/backend/model/entity"
	"github.com/tst/backend/model/report"

	"github.com/tst/backend/service"
)

func GetHistory30(ctx *fiber.Ctx) error {

	var deviceID = ctx.Query("deviceID")
	var fromDate = ctx.Query("fromDate")
	var toDate = ctx.Query("toDate")

	var DeviceReports []entity.DeviceReports
	database.DB.Debug().Table("DeviceReports").Where("DeviceID = ? and PositionStatus != 0 and ReportedOn between ? and ?", deviceID, fromDate, toDate).Order("ReportedOn").Find(&DeviceReports)

	var deviceIdleLog entity.DeviceIdleLog
	database.DB.Debug().Table("DeviceIdleLog").Where("DeviceID = ? and IdleStart >= ? and IdleEnd <= ?", deviceID, fromDate, toDate).Order("IdleStart DESC").Find(&deviceIdleLog)

	// var cache []entity.DeviceReports
	// database.DB.Debug().Table("DeviceReports").Select("top 1 *").Where("DeviceID = ? and PositionStatus != 0 and ReportedOn between ? and ?", deviceID, fromDate, toDate).Order("ReportedOn DESC").Find(&cache)

	var dev entity.Devices
	database.DB.Debug().Table("Devices").Where("DeviceID = ? ", deviceID).Find(&dev)

	var dataHistory30 []report.History30

	var history30 report.History30
	for i := 0; i < len(DeviceReports); i++ {
		history30.ReportedOn = DeviceReports[i].ReportedOn.Format("2006-01-02 15:04:05")
		history30.ReportAddress = service.NameAddress(DeviceReports[i].Latitude, DeviceReports[i].Longitude)
		history30.Speed = DeviceReports[i].Speed
		history30.Angle = enum.Angle(DeviceReports[i].Angle).String()
		history30.Latitude = DeviceReports[i].Latitude
		history30.Longitude = DeviceReports[i].Longitude

		var cache entity.DeviceReports
		database.DB.Debug().Table("DeviceReports").Select("top 1 *").Where("DeviceID = ? and PositionStatus != 0 and ReportedOn between ? and ?", deviceID, fromDate, toDate).Order("ReportedOn DESC").Where("ReportedOn <= ?",DeviceReports[i].ReportedOn).Find(&cache)

		history30.PositionStatus = enum.DeviceStatus(GetDeviceStatusReport(cache, deviceIdleLog, dev.SpeedLimit)).String()

		dataHistory30 = append(dataHistory30, history30)
	}

	return ctx.JSON(dataHistory30)

}

func GetDaily(ctx *fiber.Ctx) error {

	return nil;
}














func GetDeviceStatusReport(cache entity.DeviceReports, deviceIdleLog entity.DeviceIdleLog, speedLimit int64) int {
	report := cache
	if !report.ReportedOn.IsZero() {

		if !*report.IsWired {
			return enum.Alarm
		}
		if speedLimit > 0 && report.Speed > speedLimit {
			return enum.OverSpeed
		}
		if *report.IsAccOn {

			var DeviceIdleLog entity.DeviceIdleLog

			// Check if this is an idle stop
			if deviceIdleLog.IdleStart.Format("2006-01-02T15:04:05") <= report.ReportedOn.Format("2006-01-02T15:04:05") && deviceIdleLog.IdleEnd.Format("2006-01-02T15:04:05") >= report.ReportedOn.Format("2006-01-02T15:04:05") {
				DeviceIdleLog = deviceIdleLog
			}

			if report.Speed == 0 && !DeviceIdleLog.IdleStart.IsZero() {
				return enum.SemiStopped
			}
			return enum.Driving
		}
		return enum.Stopped
	} else {
		return enum.Stopped
	}
}
