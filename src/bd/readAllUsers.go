package bd

import (
	"context"
	"log"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("usuarios")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"email": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)

	if err != nil {
		log.Println("falla read all users" + err.Error())
		return results, false
	}

	var finded, include bool
	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		include = false
		finded, err = ConsultoRelacion(r)

		if err != nil {
			return results, false
		}

		if tipo == "new" && !finded {
			include = true
		}

		if tipo == "follow" && !finded {
			include = true
		}

		if r.UsuarioRelacionID == ID {
			include = true
		}

		if include {
			s.Pass = ""
			results = append(results, &s)
		}

	}

	err = cursor.Err()
	if err != nil {
		return results, false
	}
	cursor.Close(ctx)

	return results, true
}
