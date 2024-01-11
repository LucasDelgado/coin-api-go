package jwt

import (
	"errors"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string

var IDUsuario string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("holaQueTalComoTeVa")
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err != nil {
		return claims, false, string(""), err
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	_, exists, _ := bd.CheckBDisUser(claims.Email)

	if exists {
		Email = claims.Email
		IDUsuario = claims.ID.Hex()
	}
	return claims, exists, IDUsuario, nil
}
