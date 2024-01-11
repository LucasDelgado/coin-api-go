package bd

import (
	"context"
	"log"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(IdUser string, page int64) ([]*models.ShowTweet, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("twittors")

	var result []*models.ShowTweet

	condition := bson.M{
		"userId": IdUser,
	}

	opciones := options.Find()
	opciones.SetLimit(20)

	//para ordenar por fecha, quda comentado porque no add fecha en el models
	// opciones.SetSort(bson.D{{Key:"fecha", Value: -1}})

	opciones.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var register models.ShowTweet
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}
	return result, true
}
