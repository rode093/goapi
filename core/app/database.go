package app

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDB(dsn string) (err error) {
	log.Println("Connecting to Database Server");
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err;
	}
	DB = db;
	return;
}


