package middleware

import (
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/bd"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckBD() == 0 {
			http.Error(w, "Conexion perdida con BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
