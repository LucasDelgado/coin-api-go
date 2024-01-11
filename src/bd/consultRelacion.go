package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/LucasDelgado/coin-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("coin-app")
	col := db.Collection("relacion")

	condition := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var result models.Relacion

	fmt.Println(result)

	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println("Registro relacion no encontrado")
		return false, err
	}

	return true, nil

}
