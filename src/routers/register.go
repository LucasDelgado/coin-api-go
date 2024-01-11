package routers

import (
	//para manipular datos json
	"encoding/json"
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en body request "+err.Error(), 400)
		return
	}

	//irian las validaciones
	if len(t.Email) == 0 {
		http.Error(w, " Email de user es requerido ", 400)
		return
	}

	//chequeamos si existe el user
	_, encontrado, _ := bd.CheckBDisUser(t.Email)
	if encontrado {
		http.Error(w, " Ya existe un usuario con ese mail", 400)
		return
	}

	_, status, err := bd.AddUser(t)
	if err != nil {
		http.Error(w, " Error al intentar registar usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, " Error al intentar registar usuario status"+err.Error(), 400)
		return
	}
}
