package enum

type Angle int64

func (angle Angle) String() string {

	if angle > 337 || angle <= 22 {
		return "North (N)"
	} else if angle <= 67 {
		return "Northeast (NE)"
	} else if angle <= 112 {
		return "East (E)"
	} else if angle <= 157 {
		return "Southeast (SE)"
	} else if angle <= 202 {
		return "South (S)"
	} else if angle <= 247 {
		return "Southwest (SW)"
	} else if angle <= 292 {
		return "West (W)"
	} else if angle <= 337 {
		return "Northwest (NW)"
	}
	return "unknown"
}
