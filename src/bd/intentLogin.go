package bd

import (
	"log"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentLogin(email string, pass string) (models.User, bool) {
	usuario, exists, _ := CheckBDisUser(email)

	if !exists {
		log.Println("no existe usuario al parecer")
		return usuario, false
	}

	passToBytes := []byte(pass)
	passBD := []byte(usuario.Pass)

	err := bcrypt.CompareHashAndPassword(passBD, passToBytes)

	if err != nil {
		log.Println("falla bcrypt" + err.Error())
		return usuario, false
	}

	return usuario, true

}
