package handlers

import (
	"encoding/json"
	"fmt"
	"gormjwt/db"
	"gormjwt/models"
	"gormjwt/utils"
	"net/http"

	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// rw.WriteHeader(status)

	users := models.Users{}

	//Busco en la base de datos y tmb verifico que no haya errores en la busqueda
	if err := db.Database.Find(&users).Error; err != nil {
		fmt.Println("Error en la base de datos", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Aca verifico que no haya errores en la serializacion
	output, err := json.Marshal(&users)
	if err != nil {
		fmt.Println("error al serializar usuarios a JSON", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Si todo esta ok, escribo en el header el ok y luego con Write escribo la respuesta json en el cuerpo del http
	rw.WriteHeader(http.StatusOK)
	rw.Write(output)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if user, err := utils.GetuserById(r); err != nil {
		if err == gorm.ErrRecordNotFound {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("El usuario no existe"))
			fmt.Println("Error, no existe el usuario")
		}
		return
	} else {
		rw.WriteHeader(http.StatusAccepted)
		output, _ := json.Marshal(&user)
		rw.Write(output)
		return
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	user := models.User{}

	//Leo los datos a ingresar desde el body
	decoder := json.NewDecoder(r.Body)

	//Le ingreso el user al decoder
	if err := decoder.Decode(&user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error al decodificar el Json", err)
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "Los campos name, email y password son obligatorios")
		return
	}

	var existingUser models.User
	if err := db.Database.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		// Ya existe un usuario con el mismo correo electrónico
		rw.WriteHeader(http.StatusConflict)
		fmt.Fprintln(rw, "El correo electrónico ya está en uso por otro usuario.")
		return
	}

	if err := db.Database.Save(&user).Error; err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error al crear el usuario en la base de datos:", err)
		fmt.Fprintln(rw, "Error al crear el usuario en la base de datos:", err)
		return
	}

	// Responde con el usuario creado en formato JSON
	output, _ := json.Marshal(&user)
	fmt.Fprintln(rw, string(output))

}
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var userId int64

	if user_ant, err := utils.GetuserById(r); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("El usuario no existe"))
		return
	} else {
		userId = user_ant.ID
		user := models.User{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			rw.WriteHeader(http.StatusUnprocessableEntity)
			rw.Write([]byte("Hubo un error"))

		} else {
			user.ID = userId
			db.Database.Save(&user)
			rw.WriteHeader(http.StatusOK)
			output, _ := json.Marshal(&user)
			fmt.Fprintln(rw, string(output))
		}
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if user, err := utils.GetuserById(r); err != nil {
		fmt.Println(user)
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("El usuario no existe"))
		return
	} else {
		db.Database.Delete(&user)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("El usuario ha sido eliminado"))
	}
}
