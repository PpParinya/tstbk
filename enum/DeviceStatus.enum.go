package enum

const (
	Driving             = 1
	Stopped             = 2
	SemiStopped         = 3
	Alarm               = 4
	Unused              = 6
	LostConnection      = 7
	OverSpeed           = 9
	Manual              = 8
	NoGpsService        = 5
	IsHarshAcceleration = 10
	IsHarshBreaking     = 11
)

type DeviceStatus int64

func (s DeviceStatus) String() string {
	switch s {
	case Driving:
		return "Driving"
	case Stopped:
		return "Stopped"
	case SemiStopped:
		return "SemiStopped"
	case Alarm:
		return "Alarm"
	case Unused:
		return "Unused"
	case LostConnection:
		return "LostConnection"
	case OverSpeed:
		return "OverSpeed"
	case Manual:
		return "Manual"
	case NoGpsService:
		return "NoGpsService"
	case IsHarshAcceleration:
		return "IsHarshAcceleration"
	case IsHarshBreaking:
		return "IsHarshBreaking"
	}
	return "unknown"
}
