package handler

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/database"
	"github.com/tst/backend/enum"
	"github.com/tst/backend/model/entity"
	"github.com/tst/backend/model/response"
	"github.com/tst/backend/service"
)

const LOST_CONNECTION_TIMEOUT = 10
const TIME_UTC = 420 // + 7h

func GetAllDevices(ctx *fiber.Ctx) error {

	var UserID = ctx.Query("UserID")

	var Devices []entity.Devices
	var GetAllDevices []response.GetAllDevices
	err := database.DB.Debug().Raw("EXECUTE Web_Devices_GetListByUser  @UserID = ?", UserID).Scan(&Devices).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}
	for i := 0; i < len(Devices); i++ {
		// fmt.Println(Devices[i].DeviceID)

		var devicesByUsers response.GetAllDevices

		var users entity.Users
		database.DB.Debug().Table("Users").Where("UserID = ?", Devices[i].UserID).Find(&users)

		status := GetDeviceStatus(Devices[i], devicesByUsers)

		var deviceStatus entity.DeviceStatus
		database.DB.Debug().Table("DeviceStatus").Where("DeviceID = ?", Devices[i].DeviceID).Find(&deviceStatus)

		devicesByUsers.DeviceID = Devices[i].DeviceID
		devicesByUsers.UserID = Devices[i].UserID
		devicesByUsers.DeviceSerial = Devices[i].Serial
		devicesByUsers.PlateNumber = Devices[i].PlateNumber
		devicesByUsers.Driver = Devices[i].DriverName
		devicesByUsers.OwnerNumber = Devices[i].OwnerNumber
		devicesByUsers.Status = enum.DeviceStatus(status).String()
		devicesByUsers.StatusDetail = ""
		devicesByUsers.IconType = Devices[i].IconType
		devicesByUsers.CustomerName = users.Fullname
		devicesByUsers.IsCircuitBreakModel = Devices[i].IsCircuitBreakModel
		devicesByUsers.Remarks = Devices[i].Remarks
		devicesByUsers.TruckNumber = Devices[i].TruckNumber
		devicesByUsers.TotalMileage = (Devices[i].BaseMileage + Devices[i].AccumulatedMileage)
		devicesByUsers.DltLicense = deviceStatus.DltLicense
		devicesByUsers.DltDriverName = deviceStatus.DltDriverName
		devicesByUsers.PositionStatus = enum.DevicePositionStatus(enum.Fail).String()

		// Speed, angle, address and lat / lng
		if status != enum.Unused {

			var deviceReports entity.DeviceReports
			database.DB.Debug().Select("TOP 1 *").Table("DeviceReports").Where("DeviceID = ?", Devices[i].DeviceID).Where("(PositionStatus = ?", enum.Success).Or("PositionStatus = ?)", enum.Assisted).Order("ReportedOn DESC").Find(&deviceReports)

			devicesByUsers.Speed = deviceReports.Speed
			devicesByUsers.Angle = deviceReports.Angle
			devicesByUsers.Lat = deviceReports.Latitude
			devicesByUsers.Lng = deviceReports.Longitude
			devicesByUsers.ReportDateTime = deviceReports.ReportedOn
			devicesByUsers.IsWired = deviceReports.IsWired
			devicesByUsers.Temperature = 127
			devicesByUsers.Fuel = -1
			devicesByUsers.PositionStatus = enum.DevicePositionStatus(deviceReports.PositionStatus).String()

			if *Devices[i].HasTempSensor {

				var temperature Temperature
				database.DB.Debug().Raw("EXECUTE Web_TemperatureReports_GetLatestForDevice  @DeviceID = ?", Devices[i].DeviceID).Scan(&temperature)

				if !temperature.ReportedOn.IsZero() {
					devicesByUsers.Temperature = temperature.Temperature
				}
			}

			if *Devices[i].HasAnalogFuelSensor || *Devices[i].HasCarFuelSensor {

				var fuelSensor FuelSensor
				database.DB.Debug().Raw("EXECUTE Web_FuelReports_GetLatestForDevice  @DeviceID = ?", Devices[i].DeviceID).Scan(&fuelSensor)

				if !fuelSensor.ReportedOn.IsZero() {
					var rawFuelMax = Devices[i].RawFuelMax
					if rawFuelMax <= 0 {
						Devices[i].RawFuelMax = 1
					}

					devicesByUsers.Fuel = math.Max(Devices[i].FuelPadding, Devices[i].FuelCapacity*(float64(fuelSensor.RawFuelValue)/float64(Devices[i].RawFuelMax)))
					devicesByUsers.FuelMax = Devices[i].FuelCapacity
				}
			}

			lat, err1 := strconv.ParseFloat(devicesByUsers.Lat, 64)
			lng, err2 := strconv.ParseFloat(devicesByUsers.Lng, 64)

			if err1 != nil || err2 != nil {
				log.Println("Invalid lat/lng:", err1, err2)
			}

			// getAddress
			devicesByUsers.Address = service.NameAddress(lat, lng)
			fmt.Println(devicesByUsers.Address)

			if *Devices[i].IsCircuitBreakModel {
				devicesByUsers.IsCircuitBroken = deviceReports.IsDout1Active
			}

			var ioDescriptions DeviceIOPorts
			database.DB.Debug().Select("TOP 1 *").Table("DeviceIOPorts").Joins("JOIN DeviceIOEvents ON DeviceIOEvents.DeviceID = DeviceIOPorts.DeviceID").Where("DeviceIOPorts.DeviceID = ?", Devices[i].DeviceID).Where("DeviceIOPorts.PortType = DeviceIOEvents.PortType").Where("DeviceIOEvents.StopDate = NULL").Scan(&ioDescriptions)
			devicesByUsers.IoDescriptions = ioDescriptions.DescID

		}

		GetAllDevices = append(GetAllDevices, devicesByUsers)
	}

	return ctx.JSON(GetAllDevices)
}

func GetDeviceStatus(device entity.Devices, devicesByUsers response.GetAllDevices) enum.DeviceStatus {

	now := time.Now().UTC()

	var report CachedDeviceReportModel

	var reportCache entity.LastDeviceReport
	database.DB.Debug().Table("LastDeviceReports").Find(&reportCache, "DeviceID = ?", device.DeviceID)

	if !reportCache.ReportedOn.IsZero() {
		report.PositionStatus = enum.DevicePositionStatus(reportCache.PositionStatus).String()
		report.ReportedOn = reportCache.ReportedOn
		report.IsAccOn = *reportCache.IsAccOn
		report.Speed = reportCache.Speed
		report.IsHarshAcceleration = *reportCache.IsHarshAcceleration
		report.IsHarshBreaking = *reportCache.IsHarshBreaking
	}

	if reportCache.ReportedOn.IsZero() {
		var reportDB entity.DeviceReports
		database.DB.Debug().Select("TOP 1 *").Table("DeviceReports").Where("DeviceID = ?", device.DeviceID).Where("PositionStatus != ?", enum.Fail).Order("ReportedOn DESC").Find(&reportDB)

		if !reportDB.ReportedOn.IsZero() {
			report.PositionStatus = enum.DevicePositionStatus(reportDB.PositionStatus).String()
			report.ReportedOn = reportDB.ReportedOn
			report.IsAccOn = *reportDB.IsAccOn
			report.Speed = reportDB.Speed
		}
	}

	if device.LastHeartbeat.IsZero() || report.ReportedOn.IsZero() {
		// statusDetail = new object();
		return enum.Unused
	}

	var logEnt entity.AlarmLogs
	database.DB.Debug().Select("TOP 1 *").Table("AlarmLog").Where("DeviceID = ?", device.DeviceID).Where("AlarmLifted = ?", false).Order("AlarmOn DESC").Find(&logEnt)

	if !logEnt.AlarmOn.IsZero() {
		alarmDuration := now.Sub(logEnt.AlarmOn)
		fmt.Println(alarmDuration)

		// Determine alarm location
		var alarmLocation entity.DeviceReports
		database.DB.Debug().Select("TOP 1 *").Table("DeviceReports").Where("DeviceID = ?", device.DeviceID).Where("ReportedOn <= ?", logEnt.AlarmOn).Where("PositionStatus != ?", enum.Fail).Order("ReportedOn DESC").Find(&alarmLocation)

		if alarmLocation.ReportedOn.IsZero() {
			// statusDetail = new {
			// 	duration = alarmDuration,
			// 	alarmTime = logEnt.AlarmOn,
			// 	alarmType = logEnt.AlarmType,
			// 	alarmAddress = "No Address (Unused)",
			// 	actualStatus = DeviceStatus.Unused,
			// 	errorStatus = 0,
			// 	isAssisted = false
			// };
		} else {
			var isAssisted = false
			var actualStatus any

			if (now.Sub(device.LastHeartbeat).Minutes() + TIME_UTC) > LOST_CONNECTION_TIMEOUT {
				actualStatus = enum.LostConnection
			} else {
				var currStatus entity.DeviceReports
				database.DB.Debug().Select("TOP 1 *").Table("DeviceReports").Where("DeviceID = ?", device.DeviceID).Where("PositionStatus != ?", enum.Fail).Order("ReportedOn DESC").Find(&currStatus)

				if currStatus.ReportedOn.IsZero() {
					actualStatus = enum.Unused
				} else {
					if *currStatus.IsAccOn {
						actualStatus = enum.Driving
					} else {
						actualStatus = enum.Stopped
					}

					if currStatus.PositionStatus == enum.Assisted {
						isAssisted = true
					}
				}
			}

			var cachedAddr = ""

			// statusDetail = new {
			// 	duration = alarmDuration,
			// 	durationTime = logEnt.AlarmOn,
			// 	additionalData = logEnt.AdditionalData,
			// 	alarmType = logEnt.AlarmType,

			// 	alarmAddress= cachedAddr,
			// 	alarmLat = alarmLocation.Latitude,
			// 	alarmLng = alarmLocation.Longitude,
			// 	actualStatus = actualStatus,
			// 	errorStatus = 0,
			// 	isAssisted = isAssisted
			// };

			fmt.Println(isAssisted)
			fmt.Println(actualStatus)
			fmt.Println(cachedAddr)
		}
		return enum.Alarm
	}

	// Check last heartbeat (lost connection)
	var duration = (now.Sub(device.LastHeartbeat).Minutes() + TIME_UTC)
	// var errorStatus = 0;
	if (!*device.IsNoBatteryModel) && (duration > LOST_CONNECTION_TIMEOUT) {
		// statusDetail = new {
		// 	duration = duration.Subtract(new TimeSpan(0, LOST_CONNECTION_TIMEOUT, 0)),
		// 	isAssisted = false
		// };
		return enum.LostConnection
	}

	// Check last report (normal status)
	if report.ReportedOn.IsZero() {
		// statusDetail = new object();
		return enum.Unused
	} else {
		// duration = now - report.ReportedOn;
		// statusDetail = new {
		// 	duration = duration,
		// 	errorStatus = 0,
		// 	isAssisted = report.PositionStatus == DevicePositionStatus.Assisted
		// };

		if report.IsAccOn {

			var deviceIdleLogs entity.DeviceIdleLog
			database.DB.Debug().Select("TOP 1 *").Table("DeviceIdleLog").Where("DeviceID = ?", device.DeviceID).Where("IdleEnd IS NULL").Order("IdleStart DESC").Find(&deviceIdleLogs)
			if report.Speed == 0 && !deviceIdleLogs.IdleStart.IsZero() {
				return enum.SemiStopped
			}

			// Determine duration by finding the last ACC OFF record or registered time
			var latestAccOffReport entity.DeviceReports
			database.DB.Debug().Select("TOP 1 *").Table("DeviceReports").Where("DeviceID = ?", device.DeviceID).Where("IsAccOn = 'False'").Order("ReportedOn DESC").Find(&latestAccOffReport)

			var earliestAccOnReport entity.DeviceReports
			database.DB.Debug().Select("TOP 1 *").Table("DeviceReports").Where("DeviceID = ?", device.DeviceID).Where("IsAccOn = 'true'").Where("ReportedOn > ?", latestAccOffReport.ReportedOn).Order("ReportedOn DESC").Find(&earliestAccOnReport)

			if !earliestAccOnReport.ReportedOn.IsZero() {

				var duration = now.Sub(earliestAccOnReport.ReportedOn).Hours()

				if (duration / 24) > 1000 {
					// duration = TimeSpan.Zero;
					// errorStatus = 1;
				}
			} else if !device.RegisteredOn.IsZero() {
				// duration = now.Sub(device.ReportedOn)
			} else {
				// duration = TimeSpan.Zero;
			}

			// statusDetail = new {
			// 	duration = duration,
			// 	isAssisted = report.PositionStatus == DevicePositionStatus.Assisted,
			// 	errorStatus = errorStatus
			// };

			if report.IsHarshAcceleration {
				return enum.IsHarshAcceleration
			}
			if report.IsHarshBreaking {
				return enum.IsHarshBreaking
			}
			return enum.Driving
		}
		return enum.Stopped
	}

}

//********************************************************************** type

type CachedDeviceReportModel struct {
	Angle               int64     `json:"Angle"`
	DeviceId            int64     `json:"DeviceId"`
	ReportedOn          time.Time `json:"ReportedOn"`
	IsAccOn             bool      `json:"IsAccOn"`
	IsWired             bool      `json:"IsWired"`
	IsDout1Active       bool      `json:"IsDout1Active"`
	Speed               int64     `json:"Speed"`
	Latitude            string    `json:"Latitude"`
	Longitude           string    `json:"Longitude"`
	ReportAddress       string    `json:"ReportAddress"`
	ReportAddressEn     string    `json:"ReportAddressEn"`
	PositionStatus      string    `json:"PositionStatus"`
	IsHarshAcceleration bool      `json:"IsHarshAcceleration"`
	IsHarshBreaking     bool      `json:"IsHarshBreaking"`
}

type Temperature struct {
	ReportedOn  time.Time `json:"ReportedOn"`
	Temperature float64   `json:"Temperature"`
}

type FuelSensor struct {
	ReportedOn   time.Time `json:"ReportedOn"`
	RawFuelValue int64     `json:"RawFuelValue"`
}

type DeviceIOPorts struct {
	DeviceID int64 `json:"DeviceID"`
	PortType int64 `json:"PortType"`
	DescID   int64 `json:"DescID"`
}
