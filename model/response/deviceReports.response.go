package response

type DeviceReportsResponse struct {
	DeviceID        string `json:"DeviceID"`
	ReportedOn      string `json:"ReportedOn"`
	PositionStatus  int    `json:"PositionStatus"`
	Latitude        string `json:"Latitude"`
	Longitude       string `json:"Longitude"`
	Speed           string `json:"Speed"`
	Angle           string `json:"Angle"`
	IsWired         string `json:"IsWired"`
	IsAccOn         string `json:"IsAccOn"`
	IsDout1Active   string `json:"IsDout1Active"`
	IsSwitchOn      string `json:"IsSwitchOn"`
	ReportAddress   string `json:"ReportAddress"`
	ReportAddressEn string `json:"ReportAddressEn"`
	DltLicenseType  string `json:"DltLicenseType"`
	DltLicense      string `json:"DltLicense"`
	DltDriverName   string `json:"DltDriverName"`
}
