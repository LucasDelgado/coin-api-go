package models

type GraboTweet struct {
	UserId string `bson:"userId" json:"userId,omitempty"`
	Mess   string `bson:"mess" json:"mess,omitempty"`
}
