package response

type UserResponse struct {
	UserID       string `json:"userID"`
	Username     string `json:"username"`
	Fullname     string `json:"fullname"`
	UserType     string `json:"userType"`
	ParentUserID string `json:"parentUserID"`
}
