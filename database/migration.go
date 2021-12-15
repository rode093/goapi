package database

import (
	"learning/rest/core/app"
	"learning/rest/models"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(dsn string, schema string){
	log.Println("Connecting to DB Server...");

	db,err := app.ConnectDB(dsn, schema);
	if err != nil{
		log.Fatal("Failed to connect to DB!");
		return;
	}

	run(db);

	log.Println("Conneted to DB Server");

}

func run(db *gorm.DB){
	log.Println("Starting Migration...");
	db.AutoMigrate(&models.Conversation{});
	db.AutoMigrate(&models.Message{});

	log.Println("Migration applied.");
}

