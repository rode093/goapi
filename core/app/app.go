package app

import (
	"learning/rest/router"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Run(config map[string]string){
	if config["APP_SECRET"]=="" {
		log.Fatal("App Key is not set.");
		return;
	}

	if config["DSN"]=="" {
		log.Fatal("DB is not configured.");
		return;
	}
	db,err := connectDB(config["DSN"]);
	if(err!=nil){
		log.Fatal("Failed to connect to Database Server")
		return;
	}

	DB=db;
			log.Println("Connected to Databse Server");

	http.ListenAndServe(":8000", router.InitRouter())

}



func connectDB(dsn string) (DB *gorm.DB, err error) {
	log.Println("Connecting to Database Server");
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db,err;
}