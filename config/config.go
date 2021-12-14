package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() (config map[string]string , err error){
	
	config, err = godotenv.Read()
	if(err!= nil){
		log.Fatal("Error Loading Configuration");
		return
	}
	return
}

func Get(key string){


}

func Set(key string, val string){

}