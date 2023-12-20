package handlers

import (
	"encoding/json"
	"fmt"
	"gormjwt/db"
	"gormjwt/models"
	"net/http"
)

func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	product := models.Product{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error al decodificar el Json", err)
		return
	}

	if err := db.Database.Save(&product).Error; err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error al crear el producto en la base de datos:", err)
		fmt.Fprintln(rw, "Error al crear el producto en la base de datos:", err)
		return

	}

	output, _ := json.Marshal(&product)
	fmt.Fprintln(rw, string(output))

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []models.Product

	if err := db.Database.Find(&products).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest) // Bad request
		fmt.Println("Erro al buscar los productos")
		return
	}

	jsonResponse, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(jsonResponse)
}
