package main

import (
	"log"

	"github.com/LucasDelgado/coin-api-go/src/bd"
	"github.com/LucasDelgado/coin-api-go/src/handlers"
)

func main() {
	if bd.CheckBD() == 0 {
		log.Fatal("Sin Conexion BD")
		return
	}
	handlers.HandlersRoutes()
}
