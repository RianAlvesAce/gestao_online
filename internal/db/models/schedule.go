package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schedule struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name"`
	Schedule_in time.Time          `bson:"schedule_in"`
	CPF         string             `bson:"cpf"`
}
