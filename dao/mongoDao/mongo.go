package mongoDao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"steam-notify/config"
	"time"
)

var MongoClient *mongo.Client

func InitMongoClient(config *config.MongoDBConfig) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func getMongoClient() *mongo.Client {
	if MongoClient == nil {
		InitMongoClient(config.Config.MongoDBConfig)
	}
	return MongoClient
}

func Client(collection string) *mongo.Collection {
	return getMongoClient().Database(config.Config.MongoDBConfig.Database).Collection(collection)
}
