package jwt

import (
	"github.com/LucasDelgado/coin-api-go/src/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func JWTGenerate(t models.User) (string, error) {
	miClave := []byte("holaQueTalComoTeVa")

	payload := jwt.MapClaims{
		"email": t.Email,
		"_id":   t.ID.Hex(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
