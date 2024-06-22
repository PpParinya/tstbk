package entity

import (
	"time"
)

type Users struct {
	UserID                           string    `json:"userID" gorm:"primaryKey"`
	Username                         string    `json:"username"`
	Fullname                         string    `json:"fullname"`
	UserType                         string    `json:"userType"`
	IsEnabled                        string    `json:"isEnabled"`
	Password                         string    `json:"password"`
	FailedPasswordAttemptCount       int       `json:"failedPasswordAttemptCount"`
	FailedPasswordAttemptWindowStart time.Time `json:"failedPasswordAttemptWindowStart"`
	LastPasswordChangeOn             time.Time `json:"lastPasswordChangeOn"`
	CreatedOn                        time.Time `json:"createdOn"`
	LastLoginOn                      time.Time `json:"lastLoginOn"`
	Email                            string    `json:"email"`
	ParentUserID                     string    `json:"parentUserID"`
	Features                         int       `json:"features"`
	Preferences                      string    `json:"preferences"`
	IdleTimeout                      int       `json:"idleTimeout"`
	SpeedLimit                       int       `json:"speedLimit"`
	PeriodicReport                   *bool     `json:"periodicReport"`
	Emails                           string    `json:"emails"`
}
