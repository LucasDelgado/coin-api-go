package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LucasDelgado/coin-api-go/src/bd"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "bad req", http.StatusBadRequest)
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "bad req", http.StatusBadRequest)
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "bad req page", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	resp, status := bd.GetTweets(ID, pag)

	if !status {
		http.Error(w, "bad get tweets", 400)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)
}
