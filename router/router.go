package router

import (
	"learning/rest/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router{
	router:= mux.NewRouter();
	router.HandleFunc("/", controllers.Welcome).Methods("GET")
	 return router
}

