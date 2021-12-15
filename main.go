package main

import (
	"fmt"
	"learning/rest/config"
	"learning/rest/core/app"
	"learning/rest/database"
	"os"
)


func main()  {
	cfg, err := loadConfig()
	if err != nil{
		return;
	}
	args := os.Args[1:];

	if args[0]=="migrate"{
		database.MigrateDB(cfg["DSN"], cfg["SCHEMA"]);
		return;
	}
	
	app.Run(cfg)
	
}

func loadConfig()(c map[string]string, err error){
	//load config
	fmt.Println("Loading Config...");
	cfg, err := config.Load();
	if(err != nil){
		return nil, err;
	}
	
	return cfg, nil;
}