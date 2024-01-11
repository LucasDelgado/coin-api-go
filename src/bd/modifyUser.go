package bd

import (
	"context"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("usuarios")

	registerUpdate := make(map[string]interface{})

	// no podemos hacer checkeos dentro de una estrucutura json
	// entonces vamos armando la interfaz
	if len(u.Nombre) > 0 {
		registerUpdate["nombre"] = u.Nombre
	}

	if len(u.Email) > 0 {
		registerUpdate["email"] = u.Email
	}

	updateString := bson.M{"$set": registerUpdate}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
