package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id"`
	Login string             `bson:"login"`
	Pass  string             `bson:"pass"`
}
