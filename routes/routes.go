package routes

import (
	"gormjwt/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

}
