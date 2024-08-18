package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pacient struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `json:"name"`
	CPF      string             `json:"cpf"`
	Tel      string             `json:"tel"`
	Email    string             `json:"email"`
	Convenio int                `json:"convenio"`
}
