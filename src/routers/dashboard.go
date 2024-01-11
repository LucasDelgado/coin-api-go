package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "No hay params id", http.StatusBadRequest)
		return
	}

	userPerfil, err := bd.FindUser(ID)

	if err != nil {
		http.Error(w, "No hay params id", http.StatusBadRequest)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(userPerfil)
}
