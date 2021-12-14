package controllers

import (
	"net/http"
	"time"
)

func Welcome (response http.ResponseWriter, request *http.Request){
	
	responseData := map[string]interface{}{
		"message": "Welcome",
		"date": time.Now(),
	}
	Respond(response, &ResponseData{Status: 200, Data: responseData});
}