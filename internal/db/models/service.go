package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID          primitive.ObjectID `bson:"_id"`
	DateSV      time.Time          `bson:"dateSV"`
	Value       float64            `bson:"value"`
	PatientName string             `bson:"patientName"`
	PaymentType int                `bson:"paymentType"`
}
