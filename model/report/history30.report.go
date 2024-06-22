package report


type History30 struct {
	ReportedOn     string `json:"ReportedOn"`
	ReportAddress  string    `json:"ReportAddress"`
	Speed          int64     `json:"Speed"`
	Angle          string    `json:"Angle"`
	PositionStatus string    `json:"PositionStatus"`
	Latitude       string    `json:"Latitude"`
	Longitude      string    `json:"Longitude"`
}
