package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	el ID es de tipo primitive.ObjectID xq asi lo genera mongo. Desp ara manipularlo hay que ahcer algunas transformaciones
	el bson es el nombre de como lo guarda en la base mongo,
	el ompitempty quiere decir que si viene vacia lo omita, no lo guarde en ningun lado.
	El json: es cuando nosotros estamos devolviendo valores, es l key que va a escribir en el json q retorna.

	En definitiva, LA key del type es como manipulo el dato en el codigo, el bson es el nombre que matchea en la DB y el json con el que devolvemos valores

	BSON : Datos de entrada a la BASE
	JSON : DAtos de salida hacia el navegador.
*/

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre string             `bson:"nombre" json:"nombre,omitempty"`
	Email  string             `bson:"email" json:"email"`
	Pass   string             `bson:"pass" json:"pass,omitempty"`
}
