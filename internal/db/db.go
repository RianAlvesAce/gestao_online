package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var ServiceCollection *mongo.Collection
var PacientCollection *mongo.Collection
var ScheduleCollection *mongo.Collection 
var Ctx = context.TODO()

func InitDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))
	client, err := mongo.Connect(Ctx, clientOptions)

	if err != nil {
		log.Fatal()
	}

	err = client.Ping(Ctx, nil)

	if err != nil {
		log.Fatal()
	}

	UserCollection = client.Database("dbSGC").Collection("user")
	ServiceCollection = client.Database("dbSGC").Collection("services")
	PacientCollection = client.Database("dbSGC").Collection("pacients")
	ScheduleCollection = client.Database("dbSGC").Collection("schedules")
}