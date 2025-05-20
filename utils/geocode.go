package utils

import (
	"fmt"
	"strconv"
	"time"
)

// ReverseGeocode จำลองการแปลง lat/lng เป็นที่อยู่ (mock)
func ReverseGeocode(lat, lng float64) string {
	// ในของจริงอาจเรียก Google API หรือ Nominatim API
	time.Sleep(200 * time.Millisecond) // จำลอง delay จาก API

	// ส่งที่อยู่จำลอง
	return fmt.Sprintf("Address at %s, %s",
		strconv.FormatFloat(lat, 'f', 6, 64),
		strconv.FormatFloat(lng, 'f', 6, 64),
	)
}

// // ตัวอย่างเรียก OpenStreetMap (Nominatim)
// import (
// 	"encoding/json"
// 	"net/http"
// )

// func ReverseGeocode(lat, lng float64) string {
// 	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=%f&lon=%f", lat, lng)

// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("User-Agent", "YourAppName")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "ไม่สามารถดึงที่อยู่ได้"
// 	}
// 	defer resp.Body.Close()

// 	var result struct {
// 		DisplayName string `json:"display_name"`
// 	}
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return "ไม่สามารถอ่านข้อมูลได้"
// 	}

// 	return result.DisplayName
// }
