package request

import "time"

type Devices struct {
	UserID int64 `json:"UserID,string" gorm:"column:UserID"`
	// RouteId    int64     `json:"RouteId,string" gorm:"column:RouteId"`
	Serial     string    `json:"Serial" gorm:"column:Serial"`
	AddedOn    time.Time `json:"AddedOn" gorm:"column:AddedOn"`
	ModifiedOn time.Time `json:"ModifiedOn" gorm:"column:ModifiedOn"`
	SimNumber  string    `json:"SimNumber" gorm:"column:SimNumber"`
	// RegisteredOn            string    `json:"RegisteredOn" gorm:"column:RegisteredOn"`
	DevicePassword string `json:"DevicePassword" gorm:"column:DevicePassword"`
	OwnerNumber    string `json:"OwnerNumber" gorm:"column:OwnerNumber"`
	// LastHeartbeat           string    `json:"LastHeartbeat" gorm:"column:LastHeartbeat"`
	DriverName           string  `json:"DriverName" gorm:"column:DriverName"`
	PlateNumber          string  `json:"PlateNumber" gorm:"column:PlateNumber"`
	DeviceType           int64   `json:"DeviceType" gorm:"column:DeviceType"`
	IdleTimeout          int64   `json:"IdleTimeout" gorm:"column:IdleTimeout"`
	ReportInterval       int64   `json:"ReportInterval" gorm:"column:ReportInterval"`
	SpeedLimit           int64   `json:"SpeedLimit" gorm:"column:SpeedLimit"`
	IconType             int64   `json:"IconType" gorm:"column:IconType"`
	FuelConsumption      float64 `json:"FuelConsumption" gorm:"column:FuelConsumption"`
	AutoDisalarm         *bool   `json:"AutoDisalarm" gorm:"column:AutoDisalarm"`
	OverSpeedAlarm       *bool   `json:"OverSpeedAlarm" gorm:"column:OverSpeedAlarm"`
	IsCircuitBreakModel  *bool   `json:"IsCircuitBreakModel" gorm:"column:IsCircuitBreakModel"`
	FuelPrice            float64 `json:"FuelPrice" gorm:"column:FuelPrice"`
	HasTempSensor        *bool   `json:"HasTempSensor" gorm:"column:HasTempSensor"`
	Remarks              string  `json:"Remarks" gorm:"column:Remarks"`
	FuelCapacity         float64 `json:"FuelCapacity" gorm:"column:FuelCapacity"`
	FuelPadding          float64 `json:"FuelPadding" gorm:"column:FuelPadding"`
	BaseMileage          float64 `json:"BaseMileage" gorm:"column:BaseMileage"`
	AccumulatedMileage   float64 `json:"AccumulatedMileage" gorm:"column:AccumulatedMileage"`
	LastTriggeredMileage float64 `json:"LastTriggeredMileage" gorm:"column:LastTriggeredMileage"`
	RawFuelMax           int64   `json:"RawFuelMax" gorm:"column:RawFuelMax"`
	IsNoBatteryModel     *bool   `json:"IsNoBatteryModel" gorm:"column:IsNoBatteryModel"`
	// ModelID                 int64     `json:"ModelID" gorm:"column:ModelID"`
	// PauseRecordingToday     *bool   `json:"PauseRecordingToday" gorm:"column:PauseRecordingToday"`
	Chassis             string `json:"Chassis" gorm:"column:Chassis"`
	HasCardReader       *bool  `json:"HasCardReader" gorm:"column:HasCardReader"`
	BrandID             int64  `json:"BrandID" gorm:"column:BrandID"`
	DltRegistrationType int64  `json:"DltRegistrationType" gorm:"column:DltRegistrationType"`
	HasAnalogFuelSensor *bool  `json:"HasAnalogFuelSensor" gorm:"column:HasAnalogFuelSensor"`
	HasCarFuelSensor    *bool  `json:"HasCarFuelSensor" gorm:"column:HasCarFuelSensor"`
	NoOfSeats           int64  `json:"NoOfSeats" gorm:"column:NoOfSeats"`
	PlateProvinceID     int64  `json:"PlateProvinceID" gorm:"column:PlateProvinceID"`
	// OnBoardOverSpeedWarning *bool  `json:"OnBoardOverSpeedWarning" gorm:"column:OnBoardOverSpeedWarning"`
	// NoSwipeCardWarning      *bool  `json:"NoSwipeCardWarning" gorm:"column:NoSwipeCardWarning"`
	HarshAccelerationSpeed  int64  `json:"HarshAccelerationSpeed" gorm:"column:HarshAccelerationSpeed"`
	HarshAccelerationSecond int64  `json:"HarshAccelerationSecond" gorm:"column:HarshAccelerationSecond"`
	HarshBreakingSpeed      int64  `json:"HarshBreakingSpeed" gorm:"column:HarshBreakingSpeed"`
	HarshBreakingSecond     int64  `json:"HarshBreakingSecond" gorm:"column:HarshBreakingSecond"`
	BillingType             int64  `json:"BillingType" gorm:"column:BillingType"`
	BillingTypeChangeReason string `json:"BillingTypeChangeReason" gorm:"column:BillingTypeChangeReason"`
	OverspeedSound          int64  `json:"OverspeedSound" gorm:"column:OverspeedSound"`
	CarFuelSensorType       int64  `json:"CarFuelSensorType" gorm:"column:CarFuelSensorType"`
	TruckNumber             string `json:"TruckNumber" gorm:"column:TruckNumber"`
}
