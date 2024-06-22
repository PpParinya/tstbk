package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/config"
	"github.com/tst/backend/handler"
	"github.com/tst/backend/middleware"
)

func RouteInit(r *fiber.App) {

	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Get("queryAddress", handler.QueryAddress)
	r.Get("queryCarBrand", handler.GetCarBrand)
	r.Get("queryProvinces", handler.GetProvinces)

	r.Post("/login", handler.LoginHandler)
	r.Get("/getUsers", middleware.Auth, handler.UserHandlerGetAll)

	r.Get("/GetAllDevices", middleware.Auth, handler.GetAllDevices)

	r.Get("/GetDeviceByUser", middleware.Auth, handler.GetDeviceByUser)
	r.Post("/addDevice", middleware.Auth, handler.AddDevice)
	r.Put("/updateDevice", middleware.Auth, handler.UpdateDevice)
	r.Delete("/deleteDevice", middleware.Auth, handler.DeleteDevice)

	r.Get("/GetMarkerByUser", middleware.Auth, handler.GetMarkerByUser)
	r.Post("/addMarker", middleware.Auth, handler.AddMarker)
	r.Delete("/deleteMarker", middleware.Auth, handler.DeleteMarker)



	r.Get("/getHistory30", middleware.Auth, handler.GetHistory30)
	r.Get("/getDaily", middleware.Auth, handler.GetDaily)
}
