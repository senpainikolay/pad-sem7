package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBConnection() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+os.Getenv("MONGODB_HOST")+":"+os.Getenv("MONGODB_PORT")))
	if err != nil {
		panic(err)
	}
	return client
}
