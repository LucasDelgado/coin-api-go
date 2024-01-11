package bd

import (
	"context"
	"log"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
)

func InsertRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		log.Fatal(err.Error())
		return true, err
	}

	return true, nil
}
