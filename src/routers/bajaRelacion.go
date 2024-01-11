package routers

import (
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/jwt"
	"github.com/LucasDelgado/coin-api-go/src/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "No hay params id", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = jwt.IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.DeleteRelacion(t)

	if err != nil {
		http.Error(w, "No se pudo concretar la borrada de relacion", 400)
		return
	}

	if !status {
		http.Error(w, "no se pudo borrar relacion", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}