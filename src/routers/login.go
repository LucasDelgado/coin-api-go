package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/jwt"
	"github.com/LucasDelgado/coin-api-go/src/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, " Email o Pass invalido ", 400)
		return
	}

	_user, exists := bd.IntentLogin(t.Email, t.Pass)

	if !exists {
		http.Error(w, " Email o Pass invalido en el intento", 400)
		return
	}

	jwtKey, err := jwt.JWTGenerate(_user)

	if err != nil {
		http.Error(w, "Error al generar token", 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//grabar cookie en el front desde el back
	expTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expTime,
	})

}
