package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Database

func init() {

	cre := options.Credential{
		Username: "root",
		Password: "e9945f1586f5",
	}
	clientOptions := options.Client().ApplyURI("mongodb://172.30.179.120:27017").SetAuth(cre)
	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Errorf("mongo init err: %w", err))
	}
	MongoClient = client.Database("school_management")
}
