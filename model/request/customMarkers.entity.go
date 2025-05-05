package request

import "time"

type CustomMarkers struct {
	MarkerID          int64     `gorm:"column:MarkerID; not null; autoIncrement"`
	UserID            int64    `json:"UserID" gorm:"column:UserID"`
	Latitude          string    `json:"Latitude" gorm:"column:Latitude"`
	Longitude         string    `json:"Longitude" gorm:"column:Longitude"`
	Name              string    `json:"Name" gorm:"column:Name"`
	Description       string    `json:"Description" gorm:"column:Description"`
	CreatedBy         string    `json:"CreatedBy" gorm:"column:CreatedBy"`
	CreatedOn         time.Time `json:"CreatedOn" gorm:"column:CreatedOn"`
	Perimeter         int       `json:"Perimeter" gorm:"column:Perimeter"`
	IconType          int       `json:"IconType" gorm:"column:IconType"`
	IsCollectingPoint bool      `json:"IsCollectingPoint" gorm:"column:IsCollectingPoint"`
	IsReported        bool      `json:"IsReported" gorm:"column:IsReported"`
	SpeedLimit        int       `json:"SpeedLimit" gorm:"column:SpeedLimit"`
}
