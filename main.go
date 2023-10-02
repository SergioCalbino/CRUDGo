package main

import (
	"gormjwt/models"
	"gormjwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//Migo la db
	models.MigrateUser()

	//Creo el mux para las rutas
	router := mux.NewRouter()

	//Habilito las rutas
	routes.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))

}
