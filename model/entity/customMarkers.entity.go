package entity

import "time"

type CustomMarkers struct {
	MarkerID          int64     `json:"MarkerID"`
	UserID            int64     `json:"UserID"`
	Latitude          string    `json:"Latitude"`
	Longitude         string    `json:"Longitude"`
	Name              string    `json:"Name"`
	Description       string    `json:"Description"`
	CreatedBy         string     `json:"CreatedBy"`
	CreatedOn         time.Time `json:"CreatedOn"`
	Perimeter         int       `json:"Perimeter"`
	IconType          int       `json:"IconType"`
	IsCollectingPoint bool      `json:"IsCollectingPoint"`
	IsReported        bool      `json:"IsReported"`
	SpeedLimit        int       `json:"SpeedLimit"`
}
