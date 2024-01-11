package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/jwt"
)

func UsersList(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, " Debe enviar param page mayor a 0 "+err.Error(), 400)
		return
	}
	pag := int64(pagTemp)

	results, status := bd.ReadAllUsers(jwt.IDUsuario, pag, search, typeUser)

	if !status {
		http.Error(w, "no se pudo leer", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(results)
}
