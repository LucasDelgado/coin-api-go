package bd

import (
	"context"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("twittors")

	registro := bson.M{
		"userId": t.UserId,
		"mess":   t.Mess,
	}

	resultado, err := col.InsertOne(ctx, registro)

	if err != nil {
		return string(""), false, nil
	}

	objID, _ := resultado.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil

}
