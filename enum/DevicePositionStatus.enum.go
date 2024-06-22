package enum

const (
	Fail       = 0
	Success    = 1
	LastResult = 2
	Assisted   = 3
)

type DevicePositionStatus int64

func (s DevicePositionStatus) String() string {
	switch s {
	case Fail:
		return "Fail"
	case Success:
		return "Success"
	case LastResult:
		return "LastResult"
	case Assisted:
		return "Assisted"
	}
	return "unknown"
}
