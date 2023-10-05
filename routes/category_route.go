package routes

import (
	"gormjwt/handlers"

	"github.com/gorilla/mux"
)

func CategoryRoutes(router *mux.Router) {

	router.HandleFunc("/api/category", handlers.CreateCategory).Methods("POST")

}
