package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUser(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("usuarios")

	var perfilUser models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&perfilUser)

	if err != nil {
		fmt.Println("Registro no encontrado")
		return perfilUser, err
	}

	perfilUser.Pass = ""

	return perfilUser, nil
}
