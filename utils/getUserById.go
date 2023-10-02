package utils

import (
	"fmt"
	"gormjwt/db"
	"gormjwt/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Me creo una funcion para buscar por ID
func GetuserById(r *http.Request) (models.User, error) {
	//Leo la url con el mux.Vars
	id := mux.Vars(r)
	userId, _ := strconv.Atoi(id["id"])
	fmt.Println("User ID:", userId)
	user := models.User{}

	if err := db.Database.First(&user, userId).Error; err != nil {
		return user, err // Devolver el error si no se encuentra el usuario
	}

	return user, nil // Devolver el usuario si se encuentra
}
