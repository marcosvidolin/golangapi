package repository

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

const (
	CONNECTION_STRING = "mongodb+srv://root:root@clusterbr.r9arz.mongodb.net/questionsandanswers?retryWrites=true&w=majority"
)

func GetMontoDbClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {

		clientOptions := options.Client().ApplyURI(CONNECTION_STRING)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
