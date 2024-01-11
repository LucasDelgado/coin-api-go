package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/jwt"
	"github.com/LucasDelgado/coin-api-go/src/models"
)

func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var twt models.Tweet

	err := json.NewDecoder(r.Body).Decode(&twt)
	if err != nil {
		http.Error(w, "Error en body request "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserId: jwt.IDUsuario,
		Mess:   twt.Mess,
	}

	objStr, status, erro := bd.InsertTweet(registro)

	log.Println("SEEE" + objStr)

	if erro != nil {
		http.Error(w, "Ocurrio un error al salvar el tweet", 400)
		return
	}

	if !status {
		http.Error(w, "no se pudo salvar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
