package response

import "time"

type GetAllDevices struct {
	//DeviceStatus
	DeviceID      int64   `json:"DeviceID,string" gorm:"column:DeviceID"`
	UserID        int64   `json:"UserID,string" gorm:"column:UserID"`
	DeviceSerial  string  `json:"DeviceSerial"`
	PlateNumber   string  `json:"PlateNumber"`
	Driver        string  `json:"Driver"`
	OwnerNumber   string  `json:"OwnerNumber"`
	Status        string  `json:"Status"`
	StatusDetail  string  `json:"StatusDetail"`
	IconType      int64   `json:"IconType"`
	CustomerName  string  `json:"CustomerName"`
	Remarks       string  `json:"Remarks"`
	TruckNumber   string  `json:"TruckNumber"`
	TotalMileage  float64 `json:"TotalMileage"`
	DltLicense    string  `json:"DltLicense"`
	DltDriverName string  `json:"DltDriverName"`
	//DeviceReport
	Speed          int64     `json:"Speed"`
	Angle          int64     `json:"Angle"`
	Lat            string    `json:"Lat"`
	Lng            string    `json:"Lng"`
	ReportDateTime time.Time `json:"ReportDateTime"`
	IsWired        *bool     `json:"IsWired"`
	Temperature    float64   `json:"Temperature"`
	Fuel           float64   `json:"Fuel"`
	FuelMax        float64   `json:"FuelMax"`
	PositionStatus string    `json:"PositionStatus"`

	Address             string `json:"Address"`
	IoDescriptions      int64  `json:"IoDescriptions"`
	IsCircuitBreakModel *bool  `json:"IsCircuitBreakModel"`
	IsCircuitBroken     *bool  `json:"IsCircuitBroken"`
}
