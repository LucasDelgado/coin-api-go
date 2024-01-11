package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShowTweet struct {
	ID     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserId string             `bson:"userId" json:"userId,omitempty"`
	Mess   string             `bson:"mess" json:"mess,omitempty"`
}
