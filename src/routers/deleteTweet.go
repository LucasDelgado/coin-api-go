package routers

import (
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/jwt"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "No hay params id", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, jwt.IDUsuario)

	if err != nil {
		http.Error(w, "Hubo un error al eleiminar", 400)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
