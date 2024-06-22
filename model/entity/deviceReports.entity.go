package entity

import "time"

type DeviceReports struct {
	ReportID        int64     `json:"ReportID"`
	DeviceID        int64     `json:"DeviceID,string"`
	ReportedOn      time.Time `json:"ReportedOn"`
	PositionStatus  int64     `json:"PositionStatus"`
	Latitude        string    `json:"Latitude"`
	Longitude       string    `json:"Longitude"`
	Speed           int64     `json:"Speed"`
	Angle           int64     `json:"Angle"`
	IsWired         *bool     `json:"IsWired"`
	IsAccOn         *bool     `json:"IsAccOn"`
	IsDout1Active   *bool     `json:"IsDout1Active"`
	IsSwitchOn      *bool     `json:"IsSwitchOn"`
	ReportAddress   string    `json:"ReportAddress"`
	ReportAddressEn string    `json:"ReportAddressEn"`
	DltLicenseType  int64     `json:"DltLicenseType"`
	DltLicense      string    `json:"DltLicense"`
	DltDriverName   string    `json:"DltDriverName"`
}
