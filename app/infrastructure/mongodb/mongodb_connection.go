package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbConnection struct {
	client *mongo.Client
}

func NewMongoDbConnection() (*MongoDbConnection,error) {
	credential := options.Credential{
		Username: "admin",
		Password: "admin",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").
		SetAuth(credential)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return &MongoDbConnection{client: client}, nil
}

func (c MongoDbConnection) Disconnect() {
	err := c.client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
