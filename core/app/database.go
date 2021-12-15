package app

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
var DB *gorm.DB=nil

func ConnectDB(dsn string, s string) (db *gorm.DB, err error) {
	if DB == nil {	
		log.Println("Connecting to Database Server");
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: s + ".",
			},	
		});
		
		return db, err;
	}
	return DB, nil;
	
}
