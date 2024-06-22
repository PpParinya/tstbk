package entity

import "time"

type LastDeviceReport struct {
	DeviceID            int64     `json:"DeviceID"`
	ReportedOn          time.Time `json:"ReportedOn"`
	PositionStatus      byte      `json:"PositionStatus"`
	IsAccOn             *bool     `json:"IsAccOn"`
	Speed               int64     `json:"Speed"`
	IsHarshAcceleration *bool     `json:"IsHarshAcceleration"`
	IsHarshBreaking     *bool     `json:"IsHarshBreaking"`
}
