package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoClient *mongo.Database

const (
	LiveMongoUrl = "mongodb://live-mongo.letcode.cc:27017"
	TestMongoUrl = "mongodb://test-mongo.letcode.cc:27017"
)

func init() {
	url := TestMongoUrl
	if os.Getenv("env") == "live" {
		url = LiveMongoUrl
	}
	cre := options.Credential{
		Username: "root",
		Password: "e9945f1586f5",
	}
	clientOptions := options.Client().ApplyURI(url).SetAuth(cre)
	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Errorf("mongo init err: %w", err))
	}
	MongoClient = client.Database("school_management")
}
