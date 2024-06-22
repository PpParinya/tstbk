package entity

type DeviceStatus struct {
	DeviceId                string `json:"DeviceId"`
	Latitude                string `json:"Latitude"`
	Longitude               string `json:"Longitude"`
	PositionStatus          int64  `json:"PositionStatus"`
	Speed                   string `json:"Speed"`
	Angle                   string `json:"Angle"`
	IsWired                 string `json:"IsWired"`
	IsAccOn                 string `json:"IsAccOn"`
	DltLicense              string `json:"DltLicense"`
	DltDriverName           string `json:"DltDriverName"`
	Temperature             string `json:"Temperature"`
	RawFuelValue            string `json:"RawFuelValue"`
	IsDout1Active           string `json:"IsDout1Active"`
	Io1Active               string `json:"Io1Active"`
	Io2Active               string `json:"Io2Active"`
	Io3Active               string `json:"Io3Active"`
	Io4Active               string `json:"Io4Active"`
	DltBuzzerOnTime         string `json:"DltBuzzerOnTime"`
	LastMileageNotifiedOn   string `json:"LastMileageNotifiedOn"`
	DltLicenseType          string `json:"DltLicenseType"`
	DltOBWOffUntilEngineOff string `json:"DltOBWOffUntilEngineOff"`
	SpeedOBWWarningOn       string `json:"SpeedOBWWarningOn"`
	IsHarshAcceleration     string `json:"IsHarshAcceleration"`
	IsHarshBreaking         string `json:"IsHarshBreaking"`
}
