package repository

import (
	"context"
	"log"

	"github.com/aulang/site/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database = nil

func Database() *mongo.Database {
	if database == nil {
		clientOptions := options.Client().ApplyURI(config.Config.MongoDB.Uri)
		client, err := mongo.Connect(context.Background(), clientOptions)

		if err != nil {
			log.Fatalf("连接MongoDB失败：%v", err)
		}

		err = client.Ping(context.Background(), nil)

		if err != nil {
			log.Fatalf("连接MongoDB失败：%v", err)
		}

		database = client.Database(config.Config.MongoDB.Database)
	}

	return database
}

func Collection(name string) *mongo.Collection {
	return Database().Collection(name)
}
