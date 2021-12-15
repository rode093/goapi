package app

import (
	"learning/rest/router"
	"log"
	"net/http"
)

func Run(config map[string]string){
	if config["APP_SECRET"]=="" {
		log.Fatal("App Key is not set.");
		return;
	}

	if config["DSN"]=="" {
		log.Fatal("DB is not configured.");
		return;
	}
	db,err := ConnectDB(config["DSN"], config["SCHEMA"]);
	if(err!=nil){
		log.Fatal("Failed to connect to Database Server")
		return;
	}
	if db != nil{
		log.Println("Connected to Databse Server");
		dbConnection, _ := db.DB();
		defer dbConnection.Close();
	}
		
	http.ListenAndServe(":8000", router.InitRouter())

}
