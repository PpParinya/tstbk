package database

import (
	"fmt"

	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	// const dsn = "server=TST-PROGRAMER;user id=mon;password=8898;database=TFMS;encrypt=disable"
	// const dsn = "server=192.168.1.9;user id=mon;password=8898;database=TFMS;encrypt=disable"
	const dsn = "server=10.12.1.2;user id=TFMSServer;password=$u9efR=p;database=TFMS;encrypt=disable"
	// const dsn = "server=10.11.1.58;user id=mon;password=8898;database=TFMS;encrypt=disable"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("\n\n Cannot connect to database")
	}
	fmt.Println("\n\n Connect to database")
}
