package response

type UserResponse struct {
	UserID       string `json:"UserID"`
	Username     string `json:"Username"`
	Fullname     string `json:"Fullname"`
	UserType     string `json:"UserType"`
	ParentUserID string `json:"ParentUserID"`
}
