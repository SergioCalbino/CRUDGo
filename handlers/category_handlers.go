package handlers

import (
	"encoding/json"
	"fmt"
	"gormjwt/db"
	"gormjwt/models"
	"net/http"
)

func CreateCategory(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	category := models.Category{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&category); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "Error al decodificar el JSON")
		return

	}

	if err := db.Database.Save(&category).Error; err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error al crear la categoria en la base de datos:", err)
		fmt.Fprintln(rw, "Error al crear la categorua en la base de datos:", err)
		return
	}

	output, _ := json.Marshal(&category)
	fmt.Fprintln(rw, string(output))

}
