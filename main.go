package main

import (
	"fmt"
	"learning/rest/config"
	"learning/rest/core/app"
	"log"
)


func main()  {
	//load config
	fmt.Println("Loading Config...");
	config, err := config.Load();
	if(err != nil){
		log.Fatal("Error loading config!");
		return;
	}
	
	app.Run(config)
	
}