package response

type DeviceReportsResponse struct {
	DeviceID        string `json:"deviceID"`
	ReportedOn      string `json:"reportedOn"`
	PositionStatus  int `json:"positionStatus"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
	Speed           string `json:"speed"`
	Angle           string `json:"angle"`
	IsWired         string `json:"isWired"`
	IsAccOn         string `json:"isAccOn"`
	IsDout1Active   string `json:"isDout1Active"`
	IsSwitchOn      string `json:"isSwitchOn"`
	ReportAddress   string `json:"reportAddress"`
	ReportAddressEn string `json:"reportAddressEn"`
	DltLicenseType  string `json:"dltLicenseType"`
	DltLicense      string `json:"dltLicense"`
	DltDriverName   string `json:"dltDriverName"`
}
