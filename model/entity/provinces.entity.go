package entity

type Base_provinces struct {
	Province_id      int    `json:"Province_id"`
	Country_id       int    `json:"Country_id"`
	Province_name_en string `json:"Province_name_en"`
	Province_name_th string `json:"Province_name_th"`
	Thai_dlt_code    int    `json:"Thai_dlt_code"`
}
