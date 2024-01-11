package middleware

import (
	"net/http"

	"github.com/LucasDelgado/coin-api-go/src/jwt"
)

func JWTvalidate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, _, _, err := jwt.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Erroren el token", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
