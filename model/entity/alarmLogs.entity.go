package entity

import "time"

type AlarmLogs struct {
	AlarmID        int64     `json:"AlarmID"`
	DeviceID       int64     `json:"DeviceID"`
	AlarmOn        time.Time `json:"AlarmOn"`
	AlarmType      int64     `json:"AlarmType"`
	AlarmLifted    *bool     `json:"AlarmLifted"`
	AlarmLiftedBy  int64     `json:"AlarmLiftedBy"`
	AlarmLiftedOn  time.Time `json:"AlarmLiftedOn"`
	AdditionalData string    `json:"AdditionalData"`
}
