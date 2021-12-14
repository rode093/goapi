package controllers

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Status  int
	Data interface{}
}
func Respond(w http.ResponseWriter, response *ResponseData){
	w.WriteHeader(response.Status);
	w.Header().Add("API_VERSION", "1.0");
	json.NewEncoder(w).Encode(response.Data);
}