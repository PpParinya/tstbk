package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MYSQL *gorm.DB

func DatabaseInitMysql() {
	var err error
	// dsn := "tstadmin:8898@tcp(192.168.1.8:3306)/gps_basedata?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root@tcp(localhost)/gps_basedata?charset=utf8mb4&parseTime=True&loc=Local"
	MYSQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("\n\n Cannot connect to database")
	}
	fmt.Println("\n\n Connect to database")
}
