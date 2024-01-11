package bd

import (
	"context"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("usuarios")

	u.Pass, _ = EncryptPass(u.Pass)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
