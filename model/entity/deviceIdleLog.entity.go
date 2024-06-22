package entity

import "time"

type DeviceIdleLog struct {
	IdleID    int64     `json:"IdleID"`
	DeviceID  int64     `json:"DeviceID"`
	IdleStart time.Time `json:"IdleStart"`
	IdleEnd   time.Time `json:"IdleEnd"`
}
