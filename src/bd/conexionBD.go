package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN instancia de conexion
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://lucas:6TV9nKd5COR0N4zA@cluster0.s57nlmf.mongodb.net/?retryWrites=true&w=majority")

// ConectarBD permite conectar la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		log.Println("ERRORR BD MONGO")
		return client
	}
	log.Println("EXITOOOOOO BD MONGO")
	return client
}

// CheckBD checkeo conexion
func CheckBD() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
