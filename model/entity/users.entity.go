package entity

import (
	"time"
)

type Users struct {
	UserID                           int64     `json:"UserID,string" gorm:"column:UserID"`
	Username                         string    `gorm:"column:Username" json:"Username"`
	Fullname                         string    `gorm:"column:FullName" json:"Fullname"`
	UserType                         string    `gorm:"column:UserType" json:"UserType"`
	IsEnabled                        string    `gorm:"column:IsEnabled" json:"IsEnabled"`
	Password                         string    `gorm:"column:Password" json:"Password"`
	FailedPasswordAttemptCount       int       `gorm:"column:FailedPasswordAttemptCount" json:"FailedPasswordAttemptCount"`
	FailedPasswordAttemptWindowStart time.Time `gorm:"column:FailedPasswordAttemptWindowStart" json:"FailedPasswordAttemptWindowStart"`
	LastPasswordChangeOn             time.Time `gorm:"column:LastPasswordChangeOn" json:"LastPasswordChangeOn"`
	CreatedOn                        time.Time `gorm:"column:CreatedOn" json:"CreatedOn"`
	LastLoginOn                      time.Time `gorm:"column:LastLoginOn" json:"LastLoginOn"`
	Email                            string    `gorm:"column:Email" json:"Email"`
	ParentUserID                     *int64    `gorm:"column:ParentUserID" json:"-"`
	Features                         int       `gorm:"column:Features" json:"Features"`
	Preferences                      string    `gorm:"column:Preferences" json:"Preferences"`
	IdleTimeout                      int       `gorm:"column:IdleTimeout" json:"IdleTimeout"`
	SpeedLimit                       int       `gorm:"column:SpeedLimit" json:"SpeedLimit"`
	PeriodicReport                   *bool     `gorm:"column:PeriodicReport" json:"PeriodicReport"`
	Emails                           string    `gorm:"column:Emails" json:"Emails"`
	Children                         []*Users  `json:"children,omitempty" gorm:"-"`
}

// type Users struct {
// 	UserID                           string    `json:"userID" gorm:"primaryKey"`
// 	Username                         string    `json:"username"`
// 	Fullname                         string    `json:"fullname"`
// 	UserType                         string    `json:"userType"`
// 	IsEnabled                        string    `json:"isEnabled"`
// 	Password                         string    `json:"password"`
// 	FailedPasswordAttemptCount       int       `json:"failedPasswordAttemptCount"`
// 	FailedPasswordAttemptWindowStart time.Time `json:"failedPasswordAttemptWindowStart"`
// 	LastPasswordChangeOn             time.Time `json:"lastPasswordChangeOn"`
// 	CreatedOn                        time.Time `json:"createdOn"`
// 	LastLoginOn                      time.Time `json:"lastLoginOn"`
// 	Email                            string    `json:"email"`
// 	ParentUserID                     string    `json:"parentUserID"`
// 	Features                         int       `json:"features"`
// 	Preferences                      string    `json:"preferences"`
// 	IdleTimeout                      int       `json:"idleTimeout"`
// 	SpeedLimit                       int       `json:"speedLimit"`
// 	PeriodicReport                   *bool     `json:"periodicReport"`
// 	Emails                           string    `json:"emails"`
// }
