package database

import (
	"learning/rest/core/app"
	"log"
)

func MigrateDB(dsn string){
	log.Println("Connecting to DB Server...");

	err := app.ConnectDB(dsn);
	if err != nil{
		log.Fatal("Failed to connect to DB!");
		return;
	}

	log.Println("Conneted to DB Server");

}

func run(){
	
}

