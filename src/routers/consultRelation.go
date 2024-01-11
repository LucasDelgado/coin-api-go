package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/jwt"
	"github.com/LucasDelgado/coin-api-go/src/models"
)

func ConsultRealcion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "No hay params id", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = jwt.IDUsuario
	t.UsuarioRelacionID = ID

	var res models.ConsultRelacionResponse
	res.Status = true

	status, err := bd.ConsultoRelacion(t)

	if err != nil || !status {
		res.Status = false
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(res)
}
