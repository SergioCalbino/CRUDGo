package utils

import (
	"gormjwt/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Tiempo de expiración del token (1 hora)

	// Firma el token con la clave secreta
	tokenString, err := token.SignedString([]byte("PalabraSecreta")) // Debes convertir la cadena en un []byte
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
