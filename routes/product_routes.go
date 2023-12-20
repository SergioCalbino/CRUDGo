package routes

import (
	"gormjwt/handlers"

	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) {

	router.HandleFunc("/api/product", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product", handlers.GetProducts).Methods("GET")

}
