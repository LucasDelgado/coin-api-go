package bd

import (
	"context"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckBDisUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
