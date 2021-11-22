package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	rest2 "github.com/soalno3/Gateway/rest"
)

func main() {
	routes := rest2.InitPackage()

	//init mux routes
	router := mux.NewRouter()
	routes.InitRoutes(router)

	log.Printf("API Gateway up and running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
