package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/LucasDelgado/coin-api-go/src/middleware"
	"github.com/LucasDelgado/coin-api-go/src/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// HandlersRoutes seteo puerto
func HandlersRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/user", middleware.CheckBD(middleware.JWTvalidate(routers.Dashboard))).Methods("GET")
	router.HandleFunc("/modificar-user", middleware.CheckBD(middleware.JWTvalidate(routers.ModifiyUser))).Methods("PUT")

	router.HandleFunc("/add-twt", middleware.CheckBD(middleware.JWTvalidate(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/get-tweets", middleware.CheckBD(middleware.JWTvalidate(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/delete-tweet", middleware.CheckBD(middleware.JWTvalidate(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/up-relacion", middleware.CheckBD(middleware.JWTvalidate(routers.AltaRelacion))).Methods("GET")
	router.HandleFunc("/delete-relacion", middleware.CheckBD(middleware.JWTvalidate(routers.BajaRelacion))).Methods("DELETE")

	//TODO PROBAR ENDPOINT AJOBA
	router.HandleFunc("/consult-relation", middleware.CheckBD(middleware.JWTvalidate(routers.ConsultRealcion))).Methods("GET")

	router.HandleFunc("/users-list", middleware.CheckBD(middleware.JWTvalidate(routers.UsersList))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
