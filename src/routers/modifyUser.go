package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/jwt"
	"github.com/LucasDelgado/coin-api-go/src/models"
)

func ModifiyUser(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	status, err := bd.ModifyUser(t, jwt.IDUsuario)

	if err != nil {
		http.Error(w, "ocurrio un error al modificar el User", 400)
		return
	}

	if !status {
		http.Error(w, "ocurrio un error al modificar el User", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
